-- MySQL Script generated by MySQL Workbench
-- Mon Jun 27 11:46:50 2022
-- Model: New Model    Version: 1.0
-- MySQL Workbench Forward Engineering

SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';

-- -----------------------------------------------------
-- Schema mercadofresco
-- -----------------------------------------------------

-- -----------------------------------------------------
-- Schema mercadofresco
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `mercadofresco` DEFAULT CHARACTER SET utf8mb3 ;
USE `mercadofresco` ;

-- -----------------------------------------------------
-- Table `mercadofresco`.`buyer`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `mercadofresco`.`buyer` (
  `id` INT(11) NOT NULL  AUTO_INCREMENT,
  `card_number_id` VARCHAR(255) NOT NULL,
  `first_name` VARCHAR(255) NOT NULL,
  `last_name` VARCHAR(255) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `card_number_id_UNIQUE` (`card_number_id` ASC) VISIBLE)
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb3;


-- -----------------------------------------------------
-- Table `mercadofresco`.`localities`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `mercadofresco`.`localities` (
  `id` INT(11) NOT NULL AUTO_INCREMENT,
  `locality_name` VARCHAR(255) NOT NULL,
  `province_name` VARCHAR(255) NOT NULL,
  `country_name` VARCHAR(255) NOT NULL,
  PRIMARY KEY (`id`),
  CONSTRAINT `locality_name_unique` UNIQUE (`locality_name`))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb3;


-- -----------------------------------------------------
-- Table `mercadofresco`.`carries`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `mercadofresco`.`carries` (
  `id` INT(11) NOT NULL AUTO_INCREMENT,
  `cid` VARCHAR(255) NOT NULL,
  `company_name` VARCHAR(255) NOT NULL,
  `address` VARCHAR(255) NOT NULL,
  `telephone` VARCHAR(255) NOT NULL,
  `locality_id` INT(11) NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_carries_localities1_idx` (`locality_id` ASC) VISIBLE,
  CONSTRAINT `fk_carries_localities1`
    FOREIGN KEY (`locality_id`)
    REFERENCES `mercadofresco`.`localities` (`id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE)
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb3;


-- -----------------------------------------------------
-- Table `mercadofresco`.`warehouse`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `mercadofresco`.`warehouse` (
  `id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `address` VARCHAR(255) NOT NULL,
  `telephone` VARCHAR(255) NOT NULL,
  `warehouse_code` VARCHAR(255) NOT NULL,
  `locality_id` INT(11) NOT NULL,
  `minimum_capacity` INT NOT NULL,
  `minimum_temperature` INT (11) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `id_UNIQUE` (`id` ASC) VISIBLE,
  UNIQUE INDEX `warehouse_code_UNIQUE` (`warehouse_code` ASC) VISIBLE,
  INDEX `fk_warehouse_localities1_idx` (`locality_id` ASC) VISIBLE,
  CONSTRAINT `fk_warehouse_localities1`
    FOREIGN KEY (`locality_id`)
    REFERENCES `mercadofresco`.`localities` (`id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE)
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb3;


-- -----------------------------------------------------
-- Table `mercadofresco`.`dentists`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `mercadofresco`.`dentists` (
  `id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `card_number_id` VARCHAR(255) NOT NULL,
  `first_name` VARCHAR(255) NOT NULL,
  `last_name` VARCHAR(255) NOT NULL,
  `warehouse_id` INT(10) UNSIGNED NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `id_UNIQUE` (`id` ASC) VISIBLE,
  UNIQUE INDEX `card_number_id_UNIQUE` (`card_number_id` ASC) VISIBLE,
  INDEX `fk_dentists_warehouse1_idx` (`warehouse_id` ASC) VISIBLE,
  CONSTRAINT `fk_dentists_warehouse1`
    FOREIGN KEY (`warehouse_id`)
    REFERENCES `mercadofresco`.`warehouse` (`id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE)
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb3;


-- -----------------------------------------------------
-- Table `mercadofresco`.`products_types`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `mercadofresco`.`products_types` (
  `id` INT(11) NOT NULL AUTO_INCREMENT,
  `description` VARCHAR(255) NOT NULL,
  PRIMARY KEY (`id`))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb3;


-- -----------------------------------------------------
-- Table `mercadofresco`.`sellers`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `mercadofresco`.`sellers` (
  `id` INT(11) NOT NULL AUTO_INCREMENT,
  `cid` VARCHAR(255) NOT NULL,
  `company_name` VARCHAR(255) NOT NULL,
  `address` VARCHAR(255) NOT NULL,
  `telephone` VARCHAR(255) NOT NULL,
  `locality_id` INT(11) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `cid_UNIQUE` (`cid` ASC) VISIBLE,
  INDEX `fk_sellers_localities1_idx` (`locality_id` ASC) VISIBLE,
  CONSTRAINT `fk_sellers_localities1`
    FOREIGN KEY (`locality_id`)
    REFERENCES `mercadofresco`.`localities` (`id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE)
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb3;


-- -----------------------------------------------------
-- Table `mercadofresco`.`products`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `mercadofresco`.`products` (
  `id` INT(11) NOT NULL AUTO_INCREMENT,
  `product_code` VARCHAR(255) NOT NULL,
  `description` VARCHAR(255) NOT NULL,
  `width` DECIMAL(19,2) NOT NULL,
  `length` DECIMAL(19,2) NOT NULL,
  `height` DECIMAL(19,2) NOT NULL,
  `net_weight` DECIMAL(19,2) NOT NULL,
  `expiration_rate` DECIMAL(19,2) NOT NULL,
   `recommended_freezing_temperature` DECIMAL(19,2) NOT NULL,
  `freezing_rate` DECIMAL(19,2) NOT NULL,
  `product_type_id` INT(11) NOT NULL,
  `seller_id` INT(11) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `product_code` (`product_code` ASC) VISIBLE,
  INDEX `fk_product_sellers1_idx` (`seller_id` ASC) VISIBLE,
  INDEX `fk_product_products_types1_idx` (`product_type_id` ASC) VISIBLE,
  CONSTRAINT `fk_product_products_types1`
    FOREIGN KEY (`product_type_id`)
    REFERENCES `mercadofresco`.`products_types` (`id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE,
  CONSTRAINT `fk_product_sellers1`
    FOREIGN KEY (`seller_id`)
    REFERENCES `mercadofresco`.`sellers` (`id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE)
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8;

-- -----------------------------------------------------
-- Table `mercadofresco`.`sections`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `mercadofresco`.`sections` (
  `id` INT(11) NOT NULL AUTO_INCREMENT,
  `section_number` INT(11) NOT NULL,
  `current_capacity` INT(11) NOT NULL,
  `current_temperature` INT(11) NOT NULL,
  `maximum_capacity` INT(11) NOT NULL,
  `minimum_capacity` INT(11) NOT NULL,
  `minimum_temperature` INT(11) NOT NULL,
  `product_type_id` INT(11) NOT NULL,
  `warehouse_id` INT(10) UNSIGNED NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_section_products_types1_idx` (`product_type_id` ASC) VISIBLE,
  INDEX `fk_section_warehouse1_idx` (`warehouse_id` ASC) VISIBLE,
  UNIQUE INDEX `section_number_UNIQUE` (`section_number` ASC) VISIBLE,
  CONSTRAINT `fk_section_products_types1`
    FOREIGN KEY (`product_type_id`)
    REFERENCES `mercadofresco`.`products_types` (`id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE,
  CONSTRAINT `fk_section_warehouse1`
    FOREIGN KEY (`warehouse_id`)
    REFERENCES `mercadofresco`.`warehouse` (`id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE)
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb3;


-- -----------------------------------------------------
-- Table `mercadofresco`.`products_batches`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `mercadofresco`.`products_batches` (
  `id` INT(11) NOT NULL AUTO_INCREMENT,
  `batch_number` VARCHAR(255) NOT NULL,
  `current_quantity` INT(11) NOT NULL,
  `current_temperature` DECIMAL(19,2) NOT NULL,
  `due_date` DATETIME NOT NULL,
  `initial_quantity` INT(11) NOT NULL,
  `manufacturing_date` DATETIME NOT NULL,
  `manufacturing_hour` DATETIME NOT NULL,
  `minimum_temperature` DECIMAL(19,2) NOT NULL,
  `product_id` INT(11) NOT NULL,
  `section_id` INT(11) NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_products_batches_product1_idx` (`product_id` ASC) VISIBLE,
  INDEX `fk_products_batches_section1_idx` (`section_id` ASC) VISIBLE,
  UNIQUE INDEX `batch_number_UNIQUE`(`batch_number` ASC) VISIBLE,
  CONSTRAINT `fk_products_batches_product1`
    FOREIGN KEY (`product_id`)
    REFERENCES `mercadofresco`.`products` (`id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE,
  CONSTRAINT `fk_products_batches_section1`
    FOREIGN KEY (`section_id`)
    REFERENCES `mercadofresco`.`sections` (`id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE)
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb3;


-- -----------------------------------------------------
-- Table `mercadofresco`.`inbound_orders`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `mercadofresco`.`inbound_orders` (
  `id` INT(11) NOT NULL AUTO_INCREMENT,
  `order_date` DATETIME NOT NULL,
  `order_number` VARCHAR(255) NOT NULL,
  `dentist_id` INT(10) UNSIGNED NOT NULL,
  `product_batch_id` INT(11) NOT NULL,
  `warehouse_id` INT(10) UNSIGNED NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_inbound_details_products_batches1_idx` (`product_batch_id` ASC) VISIBLE,
  INDEX `fk_inbound_details_dentists1_idx` (`dentist_id` ASC) VISIBLE,
  INDEX `fk_inbound_details_warehouse1_idx` (`warehouse_id` ASC) VISIBLE,
  CONSTRAINT `fk_inbound_details_dentists1`
    FOREIGN KEY (`dentist_id`)
    REFERENCES `mercadofresco`.`dentists` (`id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE,
  CONSTRAINT `fk_inbound_details_products_batches1`
    FOREIGN KEY (`product_batch_id`)
    REFERENCES `mercadofresco`.`products_batches` (`id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE,
  CONSTRAINT `fk_inbound_details_warehouse1`
    FOREIGN KEY (`warehouse_id`)
    REFERENCES `mercadofresco`.`warehouse` (`id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE)
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb3;


-- -----------------------------------------------------
-- Table `mercadofresco`.`product_records`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `mercadofresco`.`product_records` (
  `id` INT(11) NOT NULL AUTO_INCREMENT,
  `last_update_date` DATETIME NOT NULL,
  `purchase_price` DECIMAL(19,2) NOT NULL,
  `sale_price` DECIMAL(19,2) NOT NULL,
  `product_id` INT(11) NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_product_records_product1_idx` (`product_id` ASC) VISIBLE,
  CONSTRAINT `fk_product_records_product1`
    FOREIGN KEY (`product_id`)
    REFERENCES `mercadofresco`.`products` (`id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE)
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb3;


-- -----------------------------------------------------
-- Table `mercadofresco`.`order_status`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `mercadofresco`.`order_status` (
  `id` INT(11) NOT NULL AUTO_INCREMENT,
  `description` VARCHAR(255) NOT NULL,
  PRIMARY KEY (`id`))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb3;


-- -----------------------------------------------------
-- Table `mercadofresco`.`purchase_orders`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `mercadofresco`.`purchase_orders` (
  `id` INT(11) NOT NULL AUTO_INCREMENT,
  `order_date` DATETIME NOT NULL,
  `order_number` VARCHAR(255) NOT NULL,  
  `tracking_code` VARCHAR(255) NOT NULL,
  `buyer_id` INT(11) NOT NULL,
  `product_record_id` INT(11) NOT NULL,
  `order_status_id` INT(11) NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_purchase_orders_product_record1_idx` (`product_record_id` ASC) VISIBLE,
  INDEX `fk_purchase_orders_buyer1_idx` (`buyer_id` ASC) VISIBLE,
  INDEX `fk_purchase_orders_order_status1_idx` (`order_status_id` ASC) VISIBLE,
  CONSTRAINT `fk_purchase_orders_buyer1`
    FOREIGN KEY (`buyer_id`)
    REFERENCES `mercadofresco`.`buyer` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_purchase_orders_product_record1`
    FOREIGN KEY (`product_record_id`)
    REFERENCES `mercadofresco`.`product_records` (`id`)   
    ON DELETE CASCADE
    ON UPDATE CASCADE,  
  CONSTRAINT `fk_purchase_orders_order_status1`
    FOREIGN KEY (`order_status_id`)
    REFERENCES `mercadofresco`.`order_status` (`id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE)
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb3;


-- -----------------------------------------------------
-- Table `mercadofresco`.`order_details`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `mercadofresco`.`order_details` (
  `id` INT(11) NOT NULL AUTO_INCREMENT,
  `clean_liness_status` VARCHAR(255) NOT NULL,
  `quantity` INT(11) NOT NULL,
  `temperature` DECIMAL(19,2) NOT NULL,
  `product_record_id` INT(11) NOT NULL,
  `purchase_order_id` INT(11) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `id_UNIQUE` (`id` ASC) VISIBLE,
  INDEX `fk_order_details_product_records1_idx` (`product_record_id` ASC) VISIBLE,
  INDEX `fk_order_details_purchase_orders1_idx` (`purchase_order_id` ASC) VISIBLE,
  CONSTRAINT `fk_order_details_product_records1`
    FOREIGN KEY (`product_record_id`)
    REFERENCES `mercadofresco`.`product_records` (`id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE,
  CONSTRAINT `fk_order_details_purchase_orders1`
    FOREIGN KEY (`purchase_order_id`)
    REFERENCES `mercadofresco`.`purchase_orders` (`id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE)
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb3;


-- -----------------------------------------------------
-- Table `mercadofresco`.`rol`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `mercadofresco`.`rol` (
  `id` INT(11) NOT NULL AUTO_INCREMENT,
  `description` VARCHAR(255) NOT NULL,
  `rol_name` VARCHAR(255) NOT NULL,
  PRIMARY KEY (`id`))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb3;


-- -----------------------------------------------------
-- Table `mercadofresco`.`users`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `mercadofresco`.`users` (
  `id` INT(11) NOT NULL AUTO_INCREMENT,
  `password` VARCHAR(255) NOT NULL,
  `username` VARCHAR(255) NOT NULL,
  PRIMARY KEY (`id`))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb3;


-- -----------------------------------------------------
-- Table `mercadofresco`.`user_rol`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `mercadofresco`.`user_rol` (
  `user_id` INT(11) NOT NULL,
  `rol_id` INT(11) NOT NULL,
  INDEX `fk_user_rol_rol1_idx` (`rol_id` ASC) VISIBLE,
  INDEX `fk_user_rol_users1_idx` (`user_id` ASC) VISIBLE,
  CONSTRAINT `fk_user_rol_rol1`
    FOREIGN KEY (`rol_id`)
    REFERENCES `mercadofresco`.`rol` (`id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE,
  CONSTRAINT `fk_user_rol_users1`
    FOREIGN KEY (`user_id`)
    REFERENCES `mercadofresco`.`users` (`id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE)
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb3;


SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;


INSERT INTO `localities` (id, locality_name, province_name, country_name) VALUES
(1,'São Paulo','SP','Brazil'),
(2,'Rio de Janeiro','RJ','Brazil'),
(3,'Nova York','NY','EUA');

INSERT INTO `sellers` VALUES
(1,'cid1','Mercado','rua 1','111',1);


INSERT INTO `products_types` VALUES
(1,'description 1');

INSERT INTO `products` VALUES
(1,'product1','teste',15.00,15.00,15.00,15.00,15.00,12.00,12.00,1,1);

INSERT INTO `product_records` VALUES
(1,'2008-11-11 13:23:44',15.50,15.50,1);

INSERT INTO `warehouse` VALUES
(1,'rua 1','11','1',1,1,2.00);

INSERT INTO `sections` VALUES
(1,1,1,1,1,1,1,1,1);

INSERT INTO `products_batches` VALUES
(1,'1',1,1.00,'2008-11-11 13:23:44',1,'2008-11-11 13:23:44','2008-11-11 13:23:44',1.00,1,1);

INSERT INTO `dentists` VALUES
(1,'1','mercado','livre',1);

INSERT INTO `inbound_orders` VALUES
(1,'2008-11-11 13:23:44','1',1,1,1);

INSERT INTO `carries` VALUES
(1,'cid1','Meli','rua 1','11-5541-4120',1);

INSERT INTO `buyer` VALUES
(1,'123456','Mercado','Livre');

INSERT INTO `order_status` VALUES
(1,'order1');

INSERT INTO `purchase_orders` VALUES
(1,'2008-11-11 13:23:44','1','1521',1,1,1);

INSERT INTO `order_details` VALUES
(1,'teste',15,15.00,1,1);


