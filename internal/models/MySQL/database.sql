CREATE TABLE `report_services` (
  `order_id` bigint NOT NULL,
  `user_id` bigint NOT NULL,
  `service_id` bigint NOT NULL,
  `service_name` varchar(255) NOT NULL,
  `cost` float NOT NULL,
  `date` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`order_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `user_balance` (
  `user_id` bigint NOT NULL,
  `balance` float NOT NULL DEFAULT '0',
  PRIMARY KEY (`user_id`),
  CONSTRAINT `balance_check` CHECK ((`balance` >= 0))
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `user_transaction` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `user_id` bigint NOT NULL,
  `type_transaction` varchar(10) NOT NULL,
  `sum_transaction` float NOT NULL,
  `comment` varchar(255) NOT NULL,
  `date` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `pagination_cost` (`id`,`sum_transaction`),
  KEY `sort_date` (`id`,`date`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `orders_balance` (
  `order_id` bigint NOT NULL,
  `user_id` bigint NOT NULL,
  `service_id` bigint NOT NULL,
  `service_name` varchar(255) NOT NULL,
  `cost` float NOT NULL,
  PRIMARY KEY (`order_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;