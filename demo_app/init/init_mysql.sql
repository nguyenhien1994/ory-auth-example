USE `demo`;

DROP TABLE IF EXISTS `user_role`;
CREATE TABLE `user_role` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `identity_id` char(36) NOT NULL,
  `role` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  CONSTRAINT `user_role_ibfk_1` FOREIGN KEY (`identity_id`) REFERENCES `identities` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

DROP TRIGGER IF EXISTS `users_role_trigger`;
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

DROP TABLE IF EXISTS `posts`;
CREATE TABLE `posts` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `identity_id` char(36) NOT NULL,
  `title` varchar(255) NOT NULL,
  `content` varchar(1024) DEFAULT NULL,
  `created_at` datetime NOT NULL DEFAULT now(),
  `updated_at` datetime NOT NULL DEFAULT now(),
  PRIMARY KEY (`id`),
  CONSTRAINT `posts_ibfk_1` FOREIGN KEY (`identity_id`) REFERENCES `identities` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- INSERT INTO `posts` (`id`, `title`, `post_details`, `owner_id`) VALUES
-- (1,	'Lorem Ipsum 1', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam sed massa neque. Ut augue magna, accumsan sed tortor et, aliquam pulvinar leo. Aenean fringilla bibendum dolor vel condimentum. Integer fringilla elit quis lorem hendrerit, eu pretium tellus convallis. Cras malesuada eget lectus id fermentum. Lorem ipsum dolor sit amet, consectetur adipiscing elit. Aliquam pellentesque consequat sollicitudin. In vel luctus massa. Aenean aliquet volutpat tempus. Nam sed massa nulla.',
-- 2),
-- (2,	'Lorem Ipsum 2', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam sed massa neque. Ut augue magna, accumsan sed tortor et, aliquam pulvinar leo. Aenean fringilla bibendum dolor vel condimentum. Integer fringilla elit quis lorem hendrerit, eu pretium tellus convallis. Cras malesuada eget lectus id fermentum. Lorem ipsum dolor sit amet, consectetur adipiscing elit. Aliquam pellentesque consequat sollicitudin. In vel luctus massa. Aenean aliquet volutpat tempus. Nam sed massa nulla.',
-- 3);
