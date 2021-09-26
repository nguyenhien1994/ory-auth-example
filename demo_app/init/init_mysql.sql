USE `demo`;

CREATE TABLE IF NOT EXISTS `user_role` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `identity_id` char(36) NOT NULL,
  `role` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  CONSTRAINT `user_role_ibfk_1` FOREIGN KEY (`identity_id`) REFERENCES `identities` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- Simple trigger to insert the data from identities to user_role
-- The first inserted user role will be `admin`, otherwise, `user`
CREATE TRIGGER users_role_trigger AFTER INSERT ON `identities` FOR EACH ROW
BEGIN
  DECLARE admin_cnt INT;
  SET admin_cnt = ( select count(*) from user_role where role = 'admin' );
  IF admin_cnt=0
  THEN
    INSERT INTO `user_role` (`identity_id`, `role`) VALUES (NEW.id, "admin");
  ELSE
    INSERT INTO `user_role` (`identity_id`, `role`) VALUES (NEW.id, "user");
  END IF;
END;
