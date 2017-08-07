-- MySQL dump 10.16  Distrib 10.2.7-MariaDB, for osx10.13 (x86_64)
--
-- Host: localhost    Database: posres
-- ------------------------------------------------------
-- Server version	10.2.7-MariaDB

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `bills`
--

DROP TABLE IF EXISTS `bills`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `bills` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `cost` double(16,4) NOT NULL DEFAULT 1.0000,
  `gross_total` double(16,4) NOT NULL DEFAULT 1.0000,
  `discount` double(16,4) NOT NULL DEFAULT 1.0000,
  `net_total` double(16,4) NOT NULL DEFAULT 1.0000,
  `balance` double(16,4) NOT NULL DEFAULT 1.0000,
  `card_paid` double(16,4) NOT NULL DEFAULT 1.0000,
  `cash_paid` double(16,4) NOT NULL DEFAULT 1.0000,
  `closed` tinyint(1) NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `customer_id` int(10) unsigned DEFAULT NULL,
  `user_id` int(10) unsigned DEFAULT NULL,
  `terminal_id` int(10) unsigned DEFAULT NULL,
  `store_id` int(10) unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `bills_customer_id_foreign` (`customer_id`),
  KEY `bills_user_id_foreign` (`user_id`),
  KEY `bills_terminal_id_foreign` (`terminal_id`),
  KEY `bills_store_id_foreign` (`store_id`),
  CONSTRAINT `bills_customer_id_foreign` FOREIGN KEY (`customer_id`) REFERENCES `customers` (`id`),
  CONSTRAINT `bills_store_id_foreign` FOREIGN KEY (`store_id`) REFERENCES `stores` (`id`),
  CONSTRAINT `bills_terminal_id_foreign` FOREIGN KEY (`terminal_id`) REFERENCES `terminals` (`id`),
  CONSTRAINT `bills_user_id_foreign` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `bills`
--

LOCK TABLES `bills` WRITE;
/*!40000 ALTER TABLE `bills` DISABLE KEYS */;
/*!40000 ALTER TABLE `bills` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `customers`
--

DROP TABLE IF EXISTS `customers`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `customers` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  `contact_no` varchar(255) DEFAULT NULL,
  `address` varchar(255) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `customers`
--

LOCK TABLES `customers` WRITE;
/*!40000 ALTER TABLE `customers` DISABLE KEYS */;
/*!40000 ALTER TABLE `customers` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `discountables`
--

DROP TABLE IF EXISTS `discountables`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `discountables` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `discountable_id` int(10) unsigned DEFAULT NULL,
  `discountable_type` varchar(255) NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `discount_id` int(10) unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `discountables_discount_id_foreign` (`discount_id`),
  CONSTRAINT `discountables_discount_id_foreign` FOREIGN KEY (`discount_id`) REFERENCES `discounts` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `discountables`
--

LOCK TABLES `discountables` WRITE;
/*!40000 ALTER TABLE `discountables` DISABLE KEYS */;
/*!40000 ALTER TABLE `discountables` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `discounts`
--

DROP TABLE IF EXISTS `discounts`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `discounts` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL,
  `from` datetime NOT NULL,
  `to` datetime NOT NULL,
  `type` enum('value','percentage') CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL,
  `value` double(8,2) NOT NULL,
  `store_id` int(10) unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `discounts_store_id_foreign` (`store_id`),
  CONSTRAINT `discounts_store_id_foreign` FOREIGN KEY (`store_id`) REFERENCES `stores` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `discounts`
--

LOCK TABLES `discounts` WRITE;
/*!40000 ALTER TABLE `discounts` DISABLE KEYS */;
/*!40000 ALTER TABLE `discounts` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `inventories`
--

DROP TABLE IF EXISTS `inventories`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `inventories` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `batch_no` varchar(255) NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `service` tinyint(1) NOT NULL DEFAULT 0,
  `product_id` int(10) unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `inventories_product_id_foreign` (`product_id`),
  CONSTRAINT `inventories_product_id_foreign` FOREIGN KEY (`product_id`) REFERENCES `products` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `inventories`
--

LOCK TABLES `inventories` WRITE;
/*!40000 ALTER TABLE `inventories` DISABLE KEYS */;
/*!40000 ALTER TABLE `inventories` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `inventory_scale`
--

DROP TABLE IF EXISTS `inventory_scale`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `inventory_scale` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `units` double(12,2) NOT NULL DEFAULT 1.00,
  `price` double(12,2) NOT NULL DEFAULT 1.00,
  `inventory_id` int(10) unsigned DEFAULT NULL,
  `scale_id` int(10) unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `inventory_scale_inventory_id_foreign` (`inventory_id`),
  KEY `inventory_scale_scale_id_foreign` (`scale_id`),
  CONSTRAINT `inventory_scale_inventory_id_foreign` FOREIGN KEY (`inventory_id`) REFERENCES `inventories` (`id`),
  CONSTRAINT `inventory_scale_scale_id_foreign` FOREIGN KEY (`scale_id`) REFERENCES `scales` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `inventory_scale`
--

LOCK TABLES `inventory_scale` WRITE;
/*!40000 ALTER TABLE `inventory_scale` DISABLE KEYS */;
/*!40000 ALTER TABLE `inventory_scale` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `migrations`
--

DROP TABLE IF EXISTS `migrations`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `migrations` (
  `id_migration` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'surrogate key',
  `name` varchar(255) DEFAULT NULL COMMENT 'migration name, unique',
  `created_at` timestamp NOT NULL DEFAULT current_timestamp() COMMENT 'date migrated or rolled back',
  `statements` longtext DEFAULT NULL COMMENT 'SQL statements for this migration',
  `rollback_statements` longtext DEFAULT NULL COMMENT 'SQL statment for rolling back migration',
  `status` enum('update','rollback') DEFAULT NULL COMMENT 'update indicates it is a normal migration while rollback means this migration is rolled back',
  PRIMARY KEY (`id_migration`)
) ENGINE=InnoDB AUTO_INCREMENT=45 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `migrations`
--

LOCK TABLES `migrations` WRITE;
/*!40000 ALTER TABLE `migrations` DISABLE KEYS */;
INSERT INTO `migrations` VALUES (1,'Stores_20170713_201929','2017-08-07 02:53:48','CREATE TABLE `stores` (\n `id` INT(10) UNSIGNED NOT NULL auto_increment ,\n `name` VARCHAR(255)  NOT NULL  ,\n `address` text    DEFAULT NULL,\n `contact` tinytext    DEFAULT NULL,\n PRIMARY KEY(  `id`),\n UNIQUE KEY `name_unique`(  `name`))ENGINE=InnoDB DEFAULT CHARSET=utf8;','DROP TABLE IF EXISTS `stores`;','rollback'),(2,'Discounts_20170713_201954','2017-08-07 02:53:48','CREATE TABLE `discounts` (\n `id` INT(10) UNSIGNED NOT NULL auto_increment ,\n `name` VARCHAR(255) COLLATE utf8_unicode_ci  NOT NULL  ,\n `from` datetime  NOT NULL  ,\n `to` datetime  NOT NULL  ,\n `type` enum(\'value\',\'percentage\') COLLATE utf8_unicode_ci  NOT NULL  ,\n `value` DOUBLE(8,2)  NOT NULL  ,\n PRIMARY KEY(  `id`),\n `store_id` int(10) UNSIGNED   DEFAULT NULL,\n KEY  `discounts_store_id_foreign`(`store_id`),\n CONSTRAINT `discounts_store_id_foreign` FOREIGN KEY (`store_id`) REFERENCES `stores` (`id`)   )ENGINE=InnoDB DEFAULT CHARSET=utf8;','DROP TABLE IF EXISTS `discounts`;','rollback'),(3,'Users_20170716_065940','2017-08-07 02:53:48','CREATE TABLE `users` (\n `id` INT(10) UNSIGNED NOT NULL auto_increment ,\n `username` VARCHAR(255)    DEFAULT NULL,\n `email` VARCHAR(255)    DEFAULT NULL,\n `password` VARCHAR(255)    DEFAULT NULL,\n `remember_token` VARCHAR(255)    DEFAULT NULL,\n `created_at` DATETIME    DEFAULT NULL,\n `updated_at` DATETIME    DEFAULT NULL,\n PRIMARY KEY(  `id`),\n UNIQUE KEY `username`(  `username`))ENGINE=InnoDB DEFAULT CHARSET=utf8;','DROP TABLE IF EXISTS `users`;','rollback'),(4,'Roles_20170716_071411','2017-08-07 02:53:48','CREATE TABLE `roles` (\n `id` INT(10) UNSIGNED NOT NULL auto_increment ,\n `name` VARCHAR(255)  NOT NULL  ,\n `display_name` VARCHAR(255)    ,\n `description` VARCHAR(255)    ,\n `created_at` DATETIME    DEFAULT NULL,\n `updated_at` DATETIME    DEFAULT NULL,\n PRIMARY KEY(  `id`))ENGINE=InnoDB DEFAULT CHARSET=utf8;','DROP TABLE IF EXISTS `roles`;','rollback'),(5,'Permissions_20170716_083154','2017-08-07 02:53:48','CREATE TABLE `permissions` (\n `id` INT(10) UNSIGNED NOT NULL auto_increment ,\n `name` VARCHAR(255)  NOT NULL  ,\n `display_name` VARCHAR(255)    ,\n `description` VARCHAR(255)    ,\n `created_at` DATETIME    DEFAULT NULL,\n `updated_at` DATETIME    DEFAULT NULL,\n PRIMARY KEY(  `id`))ENGINE=InnoDB DEFAULT CHARSET=utf8;','DROP TABLE IF EXISTS `permissions`;','rollback'),(6,'PermissionRole_20170716_083244','2017-08-07 02:53:48','CREATE TABLE `permission_role` (\n `id` INT(10) UNSIGNED NOT NULL auto_increment ,\n PRIMARY KEY(  `id`),\n UNIQUE KEY `role_permission`(  `role_id`, `permission_id`),\n `role_id` INT(10) UNSIGNED   ,\n KEY  `permission_role_role_id_foreign`(`role_id`),\n CONSTRAINT `permission_role_role_id_foreign` FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`)   ,\n `permission_id` INT(10) UNSIGNED   ,\n KEY  `permission_role_permission_id_foreign`(`permission_id`),\n CONSTRAINT `permission_role_permission_id_foreign` FOREIGN KEY (`permission_id`) REFERENCES `permissions` (`id`)   )ENGINE=InnoDB DEFAULT CHARSET=utf8;','DROP TABLE IF EXISTS `permission_role`;','rollback'),(7,'RoleUser_20170716_084548','2017-08-07 02:53:48','CREATE TABLE `role_user` (\n `id` INT(10) UNSIGNED NOT NULL auto_increment ,\n PRIMARY KEY(  `id`),\n UNIQUE KEY `role_user`(  `role_id`, `user_id`),\n `role_id` INT(10) UNSIGNED   ,\n KEY  `role_user_role_id_foreign`(`role_id`),\n CONSTRAINT `role_user_role_id_foreign` FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`)   ,\n `user_id` INT(10) UNSIGNED   ,\n KEY  `role_user_user_id_foreign`(`user_id`),\n CONSTRAINT `role_user_user_id_foreign` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)   )ENGINE=InnoDB DEFAULT CHARSET=utf8;','DROP TABLE IF EXISTS `role_user`;','rollback'),(8,'Scales_20170716_084718','2017-08-07 02:53:48','CREATE TABLE `scales` (\n `id` INT(10) UNSIGNED NOT NULL auto_increment ,\n `description` VARCHAR(255)    ,\n `unit` VARCHAR(10)    ,\n `created_at` DATETIME    DEFAULT NULL,\n `updated_at` DATETIME    DEFAULT NULL,\n PRIMARY KEY(  `id`),\n `store_id` int(10) UNSIGNED   DEFAULT NULL,\n KEY  `scales_store_id_foreign`(`store_id`),\n CONSTRAINT `scales_store_id_foreign` FOREIGN KEY (`store_id`) REFERENCES `stores` (`id`)   )ENGINE=InnoDB DEFAULT CHARSET=utf8;','DROP TABLE IF EXISTS `scales`;','rollback'),(9,'Products_20170716_085355','2017-08-07 02:53:48','CREATE TABLE `products` (\n `id` INT(10) UNSIGNED NOT NULL auto_increment ,\n `product_code` VARCHAR(255)  NOT NULL  ,\n `barcode` VARCHAR(255)    ,\n `name` VARCHAR(255)  NOT NULL  ,\n `singular` tinyint(1)    DEFAULT NULL,\n `priority` tinyint(1)    DEFAULT NULL,\n `ordering` int(10) UNSIGNED NOT NULL  ,\n `image` BLOB    DEFAULT NULL,\n `image_type` VARCHAR(20)    DEFAULT NULL,\n `service` TINYINT(1)  NOT NULL  DEFAULT 0,\n `created_at` DATETIME    DEFAULT NULL,\n `updated_at` DATETIME    DEFAULT NULL,\n PRIMARY KEY(  `id`),\n `scale_id` INT(10) UNSIGNED   DEFAULT NULL,\n KEY  `products_scale_id_foreign`(`scale_id`),\n CONSTRAINT `products_scale_id_foreign` FOREIGN KEY (`scale_id`) REFERENCES `scales` (`id`)   )ENGINE=InnoDB DEFAULT CHARSET=utf8;','DROP TABLE IF EXISTS `products`;','rollback'),(10,'StoreUser_20170716_090051','2017-08-07 02:53:48','CREATE TABLE `store_user` (\n `id` INT(10) UNSIGNED NOT NULL auto_increment ,\n PRIMARY KEY(  `id`),\n UNIQUE KEY `store_user`(  `store_id`, `user_id`),\n `store_id` INT(10) UNSIGNED   ,\n KEY  `store_user_store_id_foreign`(`store_id`),\n CONSTRAINT `store_user_store_id_foreign` FOREIGN KEY (`store_id`) REFERENCES `stores` (`id`)   ,\n `user_id` INT(10) UNSIGNED   ,\n KEY  `store_user_user_id_foreign`(`user_id`),\n CONSTRAINT `store_user_user_id_foreign` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)   )ENGINE=InnoDB DEFAULT CHARSET=utf8;','DROP TABLE IF EXISTS `store_user`;','rollback'),(11,'Discountables_20170716_090213','2017-08-07 02:53:48','CREATE TABLE `discountables` (\n `id` INT(10) UNSIGNED NOT NULL auto_increment ,\n `discountable_id` INT(10) UNSIGNED   DEFAULT NULL,\n `discountable_type` VARCHAR(255)  NOT NULL  ,\n `created_at` DATETIME    DEFAULT NULL,\n `updated_at` DATETIME    DEFAULT NULL,\n PRIMARY KEY(  `id`),\n `discount_id` INT(10) UNSIGNED   DEFAULT NULL,\n KEY  `discountables_discount_id_foreign`(`discount_id`),\n CONSTRAINT `discountables_discount_id_foreign` FOREIGN KEY (`discount_id`) REFERENCES `discounts` (`id`)   )ENGINE=InnoDB DEFAULT CHARSET=utf8;','DROP TABLE IF EXISTS `discountables`;','rollback'),(12,'ProductProduct_20170716_090500','2017-08-07 02:53:48','CREATE TABLE `product_product` (\n `id` INT(10) UNSIGNED NOT NULL auto_increment ,\n `units` DOUBLE(8,2)  NOT NULL  DEFAULT 1,\n `created_at` DATETIME    DEFAULT NULL,\n `updated_at` DATETIME    DEFAULT NULL,\n PRIMARY KEY(  `id`),\n UNIQUE KEY `parent_child`(  `parent_id`, `child_id`),\n `parent_id` INT(10) UNSIGNED   DEFAULT NULL,\n KEY  `product_product_parent_id_foreign`(`parent_id`),\n CONSTRAINT `product_product_parent_id_foreign` FOREIGN KEY (`parent_id`) REFERENCES `products` (`id`)   ,\n `child_id` INT(10) UNSIGNED   DEFAULT NULL,\n KEY  `product_product_child_id_foreign`(`child_id`),\n CONSTRAINT `product_product_child_id_foreign` FOREIGN KEY (`child_id`) REFERENCES `products` (`id`)   )ENGINE=InnoDB DEFAULT CHARSET=utf8;','DROP TABLE IF EXISTS `product_product`;','rollback'),(13,'Inventories_20170716_090814','2017-08-07 02:53:48','CREATE TABLE `inventories` (\n `id` INT(10) UNSIGNED NOT NULL auto_increment ,\n `batch_no` VARCHAR(255)  NOT NULL  ,\n `created_at` DATETIME    DEFAULT NULL,\n `updated_at` DATETIME    DEFAULT NULL,\n `service` TINYINT(1)  NOT NULL  DEFAULT 0,\n PRIMARY KEY(  `id`),\n `product_id` INT(10) UNSIGNED   DEFAULT NULL,\n KEY  `inventories_product_id_foreign`(`product_id`),\n CONSTRAINT `inventories_product_id_foreign` FOREIGN KEY (`product_id`) REFERENCES `products` (`id`)   )ENGINE=InnoDB DEFAULT CHARSET=utf8;','DROP TABLE IF EXISTS `inventories`;','rollback'),(14,'Stocks_20170716_090934','2017-08-07 02:53:48','CREATE TABLE `stocks` (\n `id` INT(10) UNSIGNED NOT NULL auto_increment ,\n `available_stock` DOUBLE(8,2)  NOT NULL  DEFAULT 1,\n PRIMARY KEY(  `id`),\n `inventory_id` INT(10) UNSIGNED   DEFAULT NULL,\n KEY  `stocks_inventory_id_foreign`(`inventory_id`),\n CONSTRAINT `stocks_inventory_id_foreign` FOREIGN KEY (`inventory_id`) REFERENCES `inventories` (`id`)   )ENGINE=InnoDB DEFAULT CHARSET=utf8;','DROP TABLE IF EXISTS `stocks`;','rollback'),(15,'InventoryScale_20170716_091044','2017-08-07 02:53:48','CREATE TABLE `inventory_scale` (\n `id` INT(10) UNSIGNED NOT NULL auto_increment ,\n `units` DOUBLE(12,2)  NOT NULL  DEFAULT 1,\n `price` DOUBLE(12,2)  NOT NULL  DEFAULT 1,\n PRIMARY KEY(  `id`),\n `inventory_id` INT(10) UNSIGNED   DEFAULT NULL,\n KEY  `inventory_scale_inventory_id_foreign`(`inventory_id`),\n CONSTRAINT `inventory_scale_inventory_id_foreign` FOREIGN KEY (`inventory_id`) REFERENCES `inventories` (`id`)   ,\n `scale_id` INT(10) UNSIGNED   DEFAULT NULL,\n KEY  `inventory_scale_scale_id_foreign`(`scale_id`),\n CONSTRAINT `inventory_scale_scale_id_foreign` FOREIGN KEY (`scale_id`) REFERENCES `scales` (`id`)   )ENGINE=InnoDB DEFAULT CHARSET=utf8;','DROP TABLE IF EXISTS `inventory_scale`;','rollback'),(16,'StockFlows_20170716_091143','2017-08-07 02:53:48','CREATE TABLE `stock_flows` (\n `id` INT(10) UNSIGNED NOT NULL auto_increment ,\n `flow` DOUBLE(4,2)  NOT NULL  DEFAULT 1,\n `stockable_id` VARCHAR(255)    DEFAULT NULL,\n `stockable_type` VARCHAR(255)    DEFAULT NULL,\n `created_at` DATETIME    DEFAULT NULL,\n `updated_at` DATETIME    DEFAULT NULL,\n PRIMARY KEY(  `id`),\n `inventory_id` INT(10) UNSIGNED   DEFAULT NULL,\n KEY  `stock_flows_inventory_id_foreign`(`inventory_id`),\n CONSTRAINT `stock_flows_inventory_id_foreign` FOREIGN KEY (`inventory_id`) REFERENCES `inventories` (`id`)   )ENGINE=InnoDB DEFAULT CHARSET=utf8;','DROP TABLE IF EXISTS `stock_flows`;','rollback'),(17,'Purchases_20170716_091344','2017-08-07 02:53:48','CREATE TABLE `purchases` (\n `id` INT(10) UNSIGNED NOT NULL auto_increment ,\n `average_cost` DOUBLE(12,2)  NOT NULL  DEFAULT 1,\n `units` DOUBLE(12,2)  NOT NULL  DEFAULT 1,\n `created_at` DATETIME    DEFAULT NULL,\n `updated_at` DATETIME    DEFAULT NULL,\n PRIMARY KEY(  `id`),\n `inventory_id` INT(10) UNSIGNED   DEFAULT NULL,\n KEY  `purchases_inventory_id_foreign`(`inventory_id`),\n CONSTRAINT `purchases_inventory_id_foreign` FOREIGN KEY (`inventory_id`) REFERENCES `inventories` (`id`)   ,\n `stock_flow_id` INT(10) UNSIGNED   DEFAULT NULL,\n KEY  `purchases_stock_flow_id_foreign`(`stock_flow_id`),\n CONSTRAINT `purchases_stock_flow_id_foreign` FOREIGN KEY (`stock_flow_id`) REFERENCES `stock_flows` (`id`)   ,\n `store_id` INT(10) UNSIGNED   DEFAULT NULL,\n KEY  `purchases_store_id_foreign`(`store_id`),\n CONSTRAINT `purchases_store_id_foreign` FOREIGN KEY (`store_id`) REFERENCES `stores` (`id`)   )ENGINE=InnoDB DEFAULT CHARSET=utf8;','DROP TABLE IF EXISTS `purchases`;','rollback'),(18,'Terminals_20170716_091545','2017-08-07 02:53:48','CREATE TABLE `terminals` (\n `id` INT(10) UNSIGNED NOT NULL auto_increment ,\n `machine` VARCHAR(255)  NOT NULL  ,\n `printer` VARCHAR(255)  NOT NULL  ,\n `balance` DOUBLE(24,8)  NOT NULL  DEFAULT 1,\n `created_at` DATETIME    DEFAULT NULL,\n `updated_at` DATETIME    DEFAULT NULL,\n PRIMARY KEY(  `id`),\n `user_id` INT(10) UNSIGNED   ,\n KEY  `terminals_user_id_foreign`(`user_id`),\n CONSTRAINT `terminals_user_id_foreign` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)   ,\n `store_id` INT(10) UNSIGNED   DEFAULT NULL,\n KEY  `terminals_store_id_foreign`(`store_id`),\n CONSTRAINT `terminals_store_id_foreign` FOREIGN KEY (`store_id`) REFERENCES `stores` (`id`)   )ENGINE=InnoDB DEFAULT CHARSET=utf8;','DROP TABLE IF EXISTS `terminals`;','rollback'),(19,'Tokens_20170716_091716','2017-08-07 02:53:48','CREATE TABLE `tokens` (\n `id` INT(10) UNSIGNED NOT NULL auto_increment ,\n `key` MEDIUMTEXT  NOT NULL  ,\n `valid_thru` DATETIME  NOT NULL  ,\n `valid` TINYINT(1)  NOT NULL  ,\n PRIMARY KEY(  `id`),\n `user_id` INT(10) UNSIGNED   ,\n KEY  `tokens_user_id_foreign`(`user_id`),\n CONSTRAINT `tokens_user_id_foreign` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)   )ENGINE=InnoDB DEFAULT CHARSET=utf8;','DROP TABLE IF EXISTS `tokens`;','rollback'),(20,'Customers_20170716_091912','2017-08-07 02:53:48','CREATE TABLE `customers` (\n `id` INT(10) UNSIGNED NOT NULL auto_increment ,\n `name` VARCHAR(255)    ,\n `contact_no` VARCHAR(255)    ,\n `address` VARCHAR(255)    ,\n `created_at` DATETIME    DEFAULT NULL,\n `updated_at` DATETIME    DEFAULT NULL,\n PRIMARY KEY(  `id`))ENGINE=InnoDB DEFAULT CHARSET=utf8;','DROP TABLE IF EXISTS `customers`;','rollback'),(21,'Bills_20170716_092032','2017-08-07 02:53:48','CREATE TABLE `bills` (\n `id` INT(10) UNSIGNED NOT NULL auto_increment ,\n `cost` DOUBLE(16,4)  NOT NULL  DEFAULT 1,\n `gross_total` DOUBLE(16,4)  NOT NULL  DEFAULT 1,\n `discount` DOUBLE(16,4)  NOT NULL  DEFAULT 1,\n `net_total` DOUBLE(16,4)  NOT NULL  DEFAULT 1,\n `balance` DOUBLE(16,4)  NOT NULL  DEFAULT 1,\n `card_paid` DOUBLE(16,4)  NOT NULL  DEFAULT 1,\n `cash_paid` DOUBLE(16,4)  NOT NULL  DEFAULT 1,\n `closed` TINYINT(1)  NOT NULL  ,\n `created_at` DATETIME    DEFAULT NULL,\n `updated_at` DATETIME    DEFAULT NULL,\n `deleted_at` DATETIME    DEFAULT NULL,\n PRIMARY KEY(  `id`),\n `customer_id` INT(10) UNSIGNED   ,\n KEY  `bills_customer_id_foreign`(`customer_id`),\n CONSTRAINT `bills_customer_id_foreign` FOREIGN KEY (`customer_id`) REFERENCES `customers` (`id`)   ,\n `user_id` INT(10) UNSIGNED   ,\n KEY  `bills_user_id_foreign`(`user_id`),\n CONSTRAINT `bills_user_id_foreign` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)   ,\n `terminal_id` INT(10) UNSIGNED   ,\n KEY  `bills_terminal_id_foreign`(`terminal_id`),\n CONSTRAINT `bills_terminal_id_foreign` FOREIGN KEY (`terminal_id`) REFERENCES `terminals` (`id`)   ,\n `store_id` INT(10) UNSIGNED   DEFAULT NULL,\n KEY  `bills_store_id_foreign`(`store_id`),\n CONSTRAINT `bills_store_id_foreign` FOREIGN KEY (`store_id`) REFERENCES `stores` (`id`)   )ENGINE=InnoDB DEFAULT CHARSET=utf8;','DROP TABLE IF EXISTS `bills`;','rollback'),(22,'Sales_20170716_092314','2017-08-07 02:53:48','CREATE TABLE `sales` (\n `id` INT(10) UNSIGNED NOT NULL auto_increment ,\n `units` DOUBLE(12,2)  NOT NULL  DEFAULT 1,\n `cost` DOUBLE(12,2)  NOT NULL  DEFAULT 1,\n `hide` TINYINT(1)  NOT NULL  ,\n `unit_price` DOUBLE(12,2)  NOT NULL  DEFAULT 1,\n `total` DOUBLE(12,2)  NOT NULL  DEFAULT 1,\n `discount` DOUBLE(12,2)  NOT NULL  DEFAULT 1,\n `amount` DOUBLE(12,2)  NOT NULL  DEFAULT 1,\n `created_at` DATETIME    DEFAULT NULL,\n `updated_at` DATETIME    DEFAULT NULL,\n PRIMARY KEY(  `id`),\n `bill_id` INT(10) UNSIGNED   ,\n KEY  `sales_bill_id_foreign`(`bill_id`),\n CONSTRAINT `sales_bill_id_foreign` FOREIGN KEY (`bill_id`) REFERENCES `bills` (`id`)   ,\n `inventory_id` INT(10) UNSIGNED   ,\n KEY  `sales_inventory_id_foreign`(`inventory_id`),\n CONSTRAINT `sales_inventory_id_foreign` FOREIGN KEY (`inventory_id`) REFERENCES `inventories` (`id`)   )ENGINE=InnoDB DEFAULT CHARSET=utf8;','DROP TABLE IF EXISTS `sales`;','rollback'),(23,'Stores_20170713_201929','2017-08-07 02:53:50','CREATE TABLE `stores` (\n `id` INT(10) UNSIGNED NOT NULL auto_increment ,\n `store_name` VARCHAR(255)  NOT NULL  ,\n `address` text    DEFAULT NULL,\n `contact` tinytext    DEFAULT NULL,\n PRIMARY KEY(  `id`),\n UNIQUE KEY `name_unique`(  `store_name`))ENGINE=InnoDB DEFAULT CHARSET=utf8;',NULL,'update'),(24,'Discounts_20170713_201954','2017-08-07 02:53:50','CREATE TABLE `discounts` (\n `id` INT(10) UNSIGNED NOT NULL auto_increment ,\n `name` VARCHAR(255) COLLATE utf8_unicode_ci  NOT NULL  ,\n `from` datetime  NOT NULL  ,\n `to` datetime  NOT NULL  ,\n `type` enum(\'value\',\'percentage\') COLLATE utf8_unicode_ci  NOT NULL  ,\n `value` DOUBLE(8,2)  NOT NULL  ,\n PRIMARY KEY(  `id`),\n `store_id` int(10) UNSIGNED   DEFAULT NULL,\n KEY  `discounts_store_id_foreign`(`store_id`),\n CONSTRAINT `discounts_store_id_foreign` FOREIGN KEY (`store_id`) REFERENCES `stores` (`id`)   )ENGINE=InnoDB DEFAULT CHARSET=utf8;',NULL,'update'),(25,'Users_20170716_065940','2017-08-07 02:53:50','CREATE TABLE `users` (\n `id` INT(10) UNSIGNED NOT NULL auto_increment ,\n `username` VARCHAR(255)    DEFAULT NULL,\n `email` VARCHAR(255)    DEFAULT NULL,\n `password` VARCHAR(255)    DEFAULT NULL,\n `remember_token` VARCHAR(255)    DEFAULT NULL,\n `created_at` DATETIME    DEFAULT NULL,\n `updated_at` DATETIME    DEFAULT NULL,\n PRIMARY KEY(  `id`),\n UNIQUE KEY `username`(  `username`))ENGINE=InnoDB DEFAULT CHARSET=utf8;',NULL,'update'),(26,'Roles_20170716_071411','2017-08-07 02:53:50','CREATE TABLE `roles` (\n `id` INT(10) UNSIGNED NOT NULL auto_increment ,\n `name` VARCHAR(255)  NOT NULL  ,\n `display_name` VARCHAR(255)    ,\n `description` VARCHAR(255)    ,\n `created_at` DATETIME    DEFAULT NULL,\n `updated_at` DATETIME    DEFAULT NULL,\n PRIMARY KEY(  `id`))ENGINE=InnoDB DEFAULT CHARSET=utf8;',NULL,'update'),(27,'Permissions_20170716_083154','2017-08-07 02:53:50','CREATE TABLE `permissions` (\n `id` INT(10) UNSIGNED NOT NULL auto_increment ,\n `name` VARCHAR(255)  NOT NULL  ,\n `display_name` VARCHAR(255)    ,\n `description` VARCHAR(255)    ,\n `created_at` DATETIME    DEFAULT NULL,\n `updated_at` DATETIME    DEFAULT NULL,\n PRIMARY KEY(  `id`))ENGINE=InnoDB DEFAULT CHARSET=utf8;',NULL,'update'),(28,'PermissionRole_20170716_083244','2017-08-07 02:53:50','CREATE TABLE `permission_role` (\n `id` INT(10) UNSIGNED NOT NULL auto_increment ,\n PRIMARY KEY(  `id`),\n UNIQUE KEY `role_permission`(  `role_id`, `permission_id`),\n `role_id` INT(10) UNSIGNED   ,\n KEY  `permission_role_role_id_foreign`(`role_id`),\n CONSTRAINT `permission_role_role_id_foreign` FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`)   ,\n `permission_id` INT(10) UNSIGNED   ,\n KEY  `permission_role_permission_id_foreign`(`permission_id`),\n CONSTRAINT `permission_role_permission_id_foreign` FOREIGN KEY (`permission_id`) REFERENCES `permissions` (`id`)   )ENGINE=InnoDB DEFAULT CHARSET=utf8;',NULL,'update'),(29,'RoleUser_20170716_084548','2017-08-07 02:53:50','CREATE TABLE `role_user` (\n `id` INT(10) UNSIGNED NOT NULL auto_increment ,\n PRIMARY KEY(  `id`),\n UNIQUE KEY `role_user`(  `role_id`, `user_id`),\n `role_id` INT(10) UNSIGNED   ,\n KEY  `role_user_role_id_foreign`(`role_id`),\n CONSTRAINT `role_user_role_id_foreign` FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`)   ,\n `user_id` INT(10) UNSIGNED   ,\n KEY  `role_user_user_id_foreign`(`user_id`),\n CONSTRAINT `role_user_user_id_foreign` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)   )ENGINE=InnoDB DEFAULT CHARSET=utf8;',NULL,'update'),(30,'Scales_20170716_084718','2017-08-07 02:53:50','CREATE TABLE `scales` (\n `id` INT(10) UNSIGNED NOT NULL auto_increment ,\n `description` VARCHAR(255)    ,\n `unit` VARCHAR(10)    ,\n `created_at` DATETIME    DEFAULT NULL,\n `updated_at` DATETIME    DEFAULT NULL,\n PRIMARY KEY(  `id`),\n `store_id` int(10) UNSIGNED   DEFAULT NULL,\n KEY  `scales_store_id_foreign`(`store_id`),\n CONSTRAINT `scales_store_id_foreign` FOREIGN KEY (`store_id`) REFERENCES `stores` (`id`)   )ENGINE=InnoDB DEFAULT CHARSET=utf8;',NULL,'update'),(31,'Products_20170716_085355','2017-08-07 02:53:50','CREATE TABLE `products` (\n `id` INT(10) UNSIGNED NOT NULL auto_increment ,\n `product_code` VARCHAR(255)  NOT NULL  ,\n `barcode` VARCHAR(255)    ,\n `name` VARCHAR(255)  NOT NULL  ,\n `singular` tinyint(1)    DEFAULT NULL,\n `priority` tinyint(1)    DEFAULT NULL,\n `ordering` int(10) UNSIGNED NOT NULL  ,\n `image` BLOB    DEFAULT NULL,\n `image_type` VARCHAR(20)    DEFAULT NULL,\n `service` TINYINT(1)  NOT NULL  DEFAULT 0,\n `created_at` DATETIME    DEFAULT NULL,\n `updated_at` DATETIME    DEFAULT NULL,\n PRIMARY KEY(  `id`),\n `scale_id` INT(10) UNSIGNED   DEFAULT NULL,\n KEY  `products_scale_id_foreign`(`scale_id`),\n CONSTRAINT `products_scale_id_foreign` FOREIGN KEY (`scale_id`) REFERENCES `scales` (`id`)   )ENGINE=InnoDB DEFAULT CHARSET=utf8;',NULL,'update'),(32,'StoreUser_20170716_090051','2017-08-07 02:53:50','CREATE TABLE `store_user` (\n `id` INT(10) UNSIGNED NOT NULL auto_increment ,\n PRIMARY KEY(  `id`),\n UNIQUE KEY `store_user`(  `stores_id`, `users_id`),\n `stores_id` INT(10) UNSIGNED   ,\n KEY  `store_user_stores_id_foreign`(`stores_id`),\n CONSTRAINT `store_user_stores_id_foreign` FOREIGN KEY (`stores_id`) REFERENCES `stores` (`id`)   ,\n `users_id` INT(10) UNSIGNED   ,\n KEY  `store_user_users_id_foreign`(`users_id`),\n CONSTRAINT `store_user_users_id_foreign` FOREIGN KEY (`users_id`) REFERENCES `users` (`id`)   )ENGINE=InnoDB DEFAULT CHARSET=utf8;',NULL,'update'),(33,'Discountables_20170716_090213','2017-08-07 02:53:50','CREATE TABLE `discountables` (\n `id` INT(10) UNSIGNED NOT NULL auto_increment ,\n `discountable_id` INT(10) UNSIGNED   DEFAULT NULL,\n `discountable_type` VARCHAR(255)  NOT NULL  ,\n `created_at` DATETIME    DEFAULT NULL,\n `updated_at` DATETIME    DEFAULT NULL,\n PRIMARY KEY(  `id`),\n `discount_id` INT(10) UNSIGNED   DEFAULT NULL,\n KEY  `discountables_discount_id_foreign`(`discount_id`),\n CONSTRAINT `discountables_discount_id_foreign` FOREIGN KEY (`discount_id`) REFERENCES `discounts` (`id`)   )ENGINE=InnoDB DEFAULT CHARSET=utf8;',NULL,'update'),(34,'ProductProduct_20170716_090500','2017-08-07 02:53:50','CREATE TABLE `product_product` (\n `id` INT(10) UNSIGNED NOT NULL auto_increment ,\n `units` DOUBLE(8,2)  NOT NULL  DEFAULT 1,\n `created_at` DATETIME    DEFAULT NULL,\n `updated_at` DATETIME    DEFAULT NULL,\n PRIMARY KEY(  `id`),\n UNIQUE KEY `parent_child`(  `parent_id`, `child_id`),\n `parent_id` INT(10) UNSIGNED   DEFAULT NULL,\n KEY  `product_product_parent_id_foreign`(`parent_id`),\n CONSTRAINT `product_product_parent_id_foreign` FOREIGN KEY (`parent_id`) REFERENCES `products` (`id`)   ,\n `child_id` INT(10) UNSIGNED   DEFAULT NULL,\n KEY  `product_product_child_id_foreign`(`child_id`),\n CONSTRAINT `product_product_child_id_foreign` FOREIGN KEY (`child_id`) REFERENCES `products` (`id`)   )ENGINE=InnoDB DEFAULT CHARSET=utf8;',NULL,'update'),(35,'Inventories_20170716_090814','2017-08-07 02:53:50','CREATE TABLE `inventories` (\n `id` INT(10) UNSIGNED NOT NULL auto_increment ,\n `batch_no` VARCHAR(255)  NOT NULL  ,\n `created_at` DATETIME    DEFAULT NULL,\n `updated_at` DATETIME    DEFAULT NULL,\n `service` TINYINT(1)  NOT NULL  DEFAULT 0,\n PRIMARY KEY(  `id`),\n `product_id` INT(10) UNSIGNED   DEFAULT NULL,\n KEY  `inventories_product_id_foreign`(`product_id`),\n CONSTRAINT `inventories_product_id_foreign` FOREIGN KEY (`product_id`) REFERENCES `products` (`id`)   )ENGINE=InnoDB DEFAULT CHARSET=utf8;',NULL,'update'),(36,'Stocks_20170716_090934','2017-08-07 02:53:50','CREATE TABLE `stocks` (\n `id` INT(10) UNSIGNED NOT NULL auto_increment ,\n `available_stock` DOUBLE(8,2)  NOT NULL  DEFAULT 1,\n PRIMARY KEY(  `id`),\n `inventory_id` INT(10) UNSIGNED   DEFAULT NULL,\n KEY  `stocks_inventory_id_foreign`(`inventory_id`),\n CONSTRAINT `stocks_inventory_id_foreign` FOREIGN KEY (`inventory_id`) REFERENCES `inventories` (`id`)   )ENGINE=InnoDB DEFAULT CHARSET=utf8;',NULL,'update'),(37,'InventoryScale_20170716_091044','2017-08-07 02:53:50','CREATE TABLE `inventory_scale` (\n `id` INT(10) UNSIGNED NOT NULL auto_increment ,\n `units` DOUBLE(12,2)  NOT NULL  DEFAULT 1,\n `price` DOUBLE(12,2)  NOT NULL  DEFAULT 1,\n PRIMARY KEY(  `id`),\n `inventory_id` INT(10) UNSIGNED   DEFAULT NULL,\n KEY  `inventory_scale_inventory_id_foreign`(`inventory_id`),\n CONSTRAINT `inventory_scale_inventory_id_foreign` FOREIGN KEY (`inventory_id`) REFERENCES `inventories` (`id`)   ,\n `scale_id` INT(10) UNSIGNED   DEFAULT NULL,\n KEY  `inventory_scale_scale_id_foreign`(`scale_id`),\n CONSTRAINT `inventory_scale_scale_id_foreign` FOREIGN KEY (`scale_id`) REFERENCES `scales` (`id`)   )ENGINE=InnoDB DEFAULT CHARSET=utf8;',NULL,'update'),(38,'StockFlows_20170716_091143','2017-08-07 02:53:50','CREATE TABLE `stock_flows` (\n `id` INT(10) UNSIGNED NOT NULL auto_increment ,\n `flow` DOUBLE(4,2)  NOT NULL  DEFAULT 1,\n `stockable_id` VARCHAR(255)    DEFAULT NULL,\n `stockable_type` VARCHAR(255)    DEFAULT NULL,\n `created_at` DATETIME    DEFAULT NULL,\n `updated_at` DATETIME    DEFAULT NULL,\n PRIMARY KEY(  `id`),\n `inventory_id` INT(10) UNSIGNED   DEFAULT NULL,\n KEY  `stock_flows_inventory_id_foreign`(`inventory_id`),\n CONSTRAINT `stock_flows_inventory_id_foreign` FOREIGN KEY (`inventory_id`) REFERENCES `inventories` (`id`)   )ENGINE=InnoDB DEFAULT CHARSET=utf8;',NULL,'update'),(39,'Purchases_20170716_091344','2017-08-07 02:53:50','CREATE TABLE `purchases` (\n `id` INT(10) UNSIGNED NOT NULL auto_increment ,\n `average_cost` DOUBLE(12,2)  NOT NULL  DEFAULT 1,\n `units` DOUBLE(12,2)  NOT NULL  DEFAULT 1,\n `created_at` DATETIME    DEFAULT NULL,\n `updated_at` DATETIME    DEFAULT NULL,\n PRIMARY KEY(  `id`),\n `inventory_id` INT(10) UNSIGNED   DEFAULT NULL,\n KEY  `purchases_inventory_id_foreign`(`inventory_id`),\n CONSTRAINT `purchases_inventory_id_foreign` FOREIGN KEY (`inventory_id`) REFERENCES `inventories` (`id`)   ,\n `stock_flow_id` INT(10) UNSIGNED   DEFAULT NULL,\n KEY  `purchases_stock_flow_id_foreign`(`stock_flow_id`),\n CONSTRAINT `purchases_stock_flow_id_foreign` FOREIGN KEY (`stock_flow_id`) REFERENCES `stock_flows` (`id`)   ,\n `store_id` INT(10) UNSIGNED   DEFAULT NULL,\n KEY  `purchases_store_id_foreign`(`store_id`),\n CONSTRAINT `purchases_store_id_foreign` FOREIGN KEY (`store_id`) REFERENCES `stores` (`id`)   )ENGINE=InnoDB DEFAULT CHARSET=utf8;',NULL,'update'),(40,'Terminals_20170716_091545','2017-08-07 02:53:50','CREATE TABLE `terminals` (\n `id` INT(10) UNSIGNED NOT NULL auto_increment ,\n `machine` VARCHAR(255)  NOT NULL  ,\n `printer` VARCHAR(255)  NOT NULL  ,\n `balance` DOUBLE(24,8)  NOT NULL  DEFAULT 1,\n `created_at` DATETIME    DEFAULT NULL,\n `updated_at` DATETIME    DEFAULT NULL,\n PRIMARY KEY(  `id`),\n `user_id` INT(10) UNSIGNED   ,\n KEY  `terminals_user_id_foreign`(`user_id`),\n CONSTRAINT `terminals_user_id_foreign` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)   ,\n `store_id` INT(10) UNSIGNED   DEFAULT NULL,\n KEY  `terminals_store_id_foreign`(`store_id`),\n CONSTRAINT `terminals_store_id_foreign` FOREIGN KEY (`store_id`) REFERENCES `stores` (`id`)   )ENGINE=InnoDB DEFAULT CHARSET=utf8;',NULL,'update'),(41,'Tokens_20170716_091716','2017-08-07 02:53:51','CREATE TABLE `tokens` (\n `id` INT(10) UNSIGNED NOT NULL auto_increment ,\n `key` MEDIUMTEXT  NOT NULL  ,\n `valid_thru` DATETIME  NOT NULL  ,\n `valid` TINYINT(1)  NOT NULL  ,\n PRIMARY KEY(  `id`),\n `user_id` INT(10) UNSIGNED   ,\n KEY  `tokens_user_id_foreign`(`user_id`),\n CONSTRAINT `tokens_user_id_foreign` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)   )ENGINE=InnoDB DEFAULT CHARSET=utf8;',NULL,'update'),(42,'Customers_20170716_091912','2017-08-07 02:53:51','CREATE TABLE `customers` (\n `id` INT(10) UNSIGNED NOT NULL auto_increment ,\n `name` VARCHAR(255)    ,\n `contact_no` VARCHAR(255)    ,\n `address` VARCHAR(255)    ,\n `created_at` DATETIME    DEFAULT NULL,\n `updated_at` DATETIME    DEFAULT NULL,\n PRIMARY KEY(  `id`))ENGINE=InnoDB DEFAULT CHARSET=utf8;',NULL,'update'),(43,'Bills_20170716_092032','2017-08-07 02:53:51','CREATE TABLE `bills` (\n `id` INT(10) UNSIGNED NOT NULL auto_increment ,\n `cost` DOUBLE(16,4)  NOT NULL  DEFAULT 1,\n `gross_total` DOUBLE(16,4)  NOT NULL  DEFAULT 1,\n `discount` DOUBLE(16,4)  NOT NULL  DEFAULT 1,\n `net_total` DOUBLE(16,4)  NOT NULL  DEFAULT 1,\n `balance` DOUBLE(16,4)  NOT NULL  DEFAULT 1,\n `card_paid` DOUBLE(16,4)  NOT NULL  DEFAULT 1,\n `cash_paid` DOUBLE(16,4)  NOT NULL  DEFAULT 1,\n `closed` TINYINT(1)  NOT NULL  ,\n `created_at` DATETIME    DEFAULT NULL,\n `updated_at` DATETIME    DEFAULT NULL,\n `deleted_at` DATETIME    DEFAULT NULL,\n PRIMARY KEY(  `id`),\n `customer_id` INT(10) UNSIGNED   ,\n KEY  `bills_customer_id_foreign`(`customer_id`),\n CONSTRAINT `bills_customer_id_foreign` FOREIGN KEY (`customer_id`) REFERENCES `customers` (`id`)   ,\n `user_id` INT(10) UNSIGNED   ,\n KEY  `bills_user_id_foreign`(`user_id`),\n CONSTRAINT `bills_user_id_foreign` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)   ,\n `terminal_id` INT(10) UNSIGNED   ,\n KEY  `bills_terminal_id_foreign`(`terminal_id`),\n CONSTRAINT `bills_terminal_id_foreign` FOREIGN KEY (`terminal_id`) REFERENCES `terminals` (`id`)   ,\n `store_id` INT(10) UNSIGNED   DEFAULT NULL,\n KEY  `bills_store_id_foreign`(`store_id`),\n CONSTRAINT `bills_store_id_foreign` FOREIGN KEY (`store_id`) REFERENCES `stores` (`id`)   )ENGINE=InnoDB DEFAULT CHARSET=utf8;',NULL,'update'),(44,'Sales_20170716_092314','2017-08-07 02:53:51','CREATE TABLE `sales` (\n `id` INT(10) UNSIGNED NOT NULL auto_increment ,\n `units` DOUBLE(12,2)  NOT NULL  DEFAULT 1,\n `cost` DOUBLE(12,2)  NOT NULL  DEFAULT 1,\n `hide` TINYINT(1)  NOT NULL  ,\n `unit_price` DOUBLE(12,2)  NOT NULL  DEFAULT 1,\n `total` DOUBLE(12,2)  NOT NULL  DEFAULT 1,\n `discount` DOUBLE(12,2)  NOT NULL  DEFAULT 1,\n `amount` DOUBLE(12,2)  NOT NULL  DEFAULT 1,\n `created_at` DATETIME    DEFAULT NULL,\n `updated_at` DATETIME    DEFAULT NULL,\n `deleted_at` DATETIME    DEFAULT NULL,\n PRIMARY KEY(  `id`),\n `bill_id` INT(10) UNSIGNED   ,\n KEY  `sales_bill_id_foreign`(`bill_id`),\n CONSTRAINT `sales_bill_id_foreign` FOREIGN KEY (`bill_id`) REFERENCES `bills` (`id`)   ,\n `inventory_id` INT(10) UNSIGNED   ,\n KEY  `sales_inventory_id_foreign`(`inventory_id`),\n CONSTRAINT `sales_inventory_id_foreign` FOREIGN KEY (`inventory_id`) REFERENCES `inventories` (`id`)   )ENGINE=InnoDB DEFAULT CHARSET=utf8;',NULL,'update');
/*!40000 ALTER TABLE `migrations` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `password_resets`
--

DROP TABLE IF EXISTS `password_resets`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `password_resets` (
  `email` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `token` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  KEY `password_resets_email_index` (`email`),
  KEY `password_resets_token_index` (`token`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `password_resets`
--

LOCK TABLES `password_resets` WRITE;
/*!40000 ALTER TABLE `password_resets` DISABLE KEYS */;
/*!40000 ALTER TABLE `password_resets` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `permission_role`
--

DROP TABLE IF EXISTS `permission_role`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `permission_role` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `role_id` int(10) unsigned DEFAULT NULL,
  `permission_id` int(10) unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `role_permission` (`role_id`,`permission_id`),
  KEY `permission_role_role_id_foreign` (`role_id`),
  KEY `permission_role_permission_id_foreign` (`permission_id`),
  CONSTRAINT `permission_role_permission_id_foreign` FOREIGN KEY (`permission_id`) REFERENCES `permissions` (`id`),
  CONSTRAINT `permission_role_role_id_foreign` FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `permission_role`
--

LOCK TABLES `permission_role` WRITE;
/*!40000 ALTER TABLE `permission_role` DISABLE KEYS */;
/*!40000 ALTER TABLE `permission_role` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `permissions`
--

DROP TABLE IF EXISTS `permissions`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `permissions` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `display_name` varchar(255) DEFAULT NULL,
  `description` varchar(255) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `permissions`
--

LOCK TABLES `permissions` WRITE;
/*!40000 ALTER TABLE `permissions` DISABLE KEYS */;
/*!40000 ALTER TABLE `permissions` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `product_product`
--

DROP TABLE IF EXISTS `product_product`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `product_product` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `units` double(8,2) NOT NULL DEFAULT 1.00,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `parent_id` int(10) unsigned DEFAULT NULL,
  `child_id` int(10) unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `parent_child` (`parent_id`,`child_id`),
  KEY `product_product_parent_id_foreign` (`parent_id`),
  KEY `product_product_child_id_foreign` (`child_id`),
  CONSTRAINT `product_product_child_id_foreign` FOREIGN KEY (`child_id`) REFERENCES `products` (`id`),
  CONSTRAINT `product_product_parent_id_foreign` FOREIGN KEY (`parent_id`) REFERENCES `products` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `product_product`
--

LOCK TABLES `product_product` WRITE;
/*!40000 ALTER TABLE `product_product` DISABLE KEYS */;
/*!40000 ALTER TABLE `product_product` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `products`
--

DROP TABLE IF EXISTS `products`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `products` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `product_code` varchar(255) NOT NULL,
  `barcode` varchar(255) DEFAULT NULL,
  `name` varchar(255) NOT NULL,
  `singular` tinyint(1) DEFAULT NULL,
  `priority` tinyint(1) DEFAULT NULL,
  `ordering` int(10) unsigned NOT NULL,
  `image` blob DEFAULT NULL,
  `image_type` varchar(20) DEFAULT NULL,
  `service` tinyint(1) NOT NULL DEFAULT 0,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `scale_id` int(10) unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `products_scale_id_foreign` (`scale_id`),
  CONSTRAINT `products_scale_id_foreign` FOREIGN KEY (`scale_id`) REFERENCES `scales` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `products`
--

LOCK TABLES `products` WRITE;
/*!40000 ALTER TABLE `products` DISABLE KEYS */;
/*!40000 ALTER TABLE `products` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `purchases`
--

DROP TABLE IF EXISTS `purchases`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `purchases` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `average_cost` double(12,2) NOT NULL DEFAULT 1.00,
  `units` double(12,2) NOT NULL DEFAULT 1.00,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `inventory_id` int(10) unsigned DEFAULT NULL,
  `stock_flow_id` int(10) unsigned DEFAULT NULL,
  `store_id` int(10) unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `purchases_inventory_id_foreign` (`inventory_id`),
  KEY `purchases_stock_flow_id_foreign` (`stock_flow_id`),
  KEY `purchases_store_id_foreign` (`store_id`),
  CONSTRAINT `purchases_inventory_id_foreign` FOREIGN KEY (`inventory_id`) REFERENCES `inventories` (`id`),
  CONSTRAINT `purchases_stock_flow_id_foreign` FOREIGN KEY (`stock_flow_id`) REFERENCES `stock_flows` (`id`),
  CONSTRAINT `purchases_store_id_foreign` FOREIGN KEY (`store_id`) REFERENCES `stores` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `purchases`
--

LOCK TABLES `purchases` WRITE;
/*!40000 ALTER TABLE `purchases` DISABLE KEYS */;
/*!40000 ALTER TABLE `purchases` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `role_user`
--

DROP TABLE IF EXISTS `role_user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `role_user` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `role_id` int(10) unsigned DEFAULT NULL,
  `user_id` int(10) unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `role_user` (`role_id`,`user_id`),
  KEY `role_user_role_id_foreign` (`role_id`),
  KEY `role_user_user_id_foreign` (`user_id`),
  CONSTRAINT `role_user_role_id_foreign` FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`),
  CONSTRAINT `role_user_user_id_foreign` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `role_user`
--

LOCK TABLES `role_user` WRITE;
/*!40000 ALTER TABLE `role_user` DISABLE KEYS */;
/*!40000 ALTER TABLE `role_user` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `roles`
--

DROP TABLE IF EXISTS `roles`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `roles` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `display_name` varchar(255) DEFAULT NULL,
  `description` varchar(255) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `roles`
--

LOCK TABLES `roles` WRITE;
/*!40000 ALTER TABLE `roles` DISABLE KEYS */;
/*!40000 ALTER TABLE `roles` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sales`
--

DROP TABLE IF EXISTS `sales`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sales` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `units` double(12,2) NOT NULL DEFAULT 1.00,
  `cost` double(12,2) NOT NULL DEFAULT 1.00,
  `hide` tinyint(1) NOT NULL,
  `unit_price` double(12,2) NOT NULL DEFAULT 1.00,
  `total` double(12,2) NOT NULL DEFAULT 1.00,
  `discount` double(12,2) NOT NULL DEFAULT 1.00,
  `amount` double(12,2) NOT NULL DEFAULT 1.00,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `bill_id` int(10) unsigned DEFAULT NULL,
  `inventory_id` int(10) unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `sales_bill_id_foreign` (`bill_id`),
  KEY `sales_inventory_id_foreign` (`inventory_id`),
  CONSTRAINT `sales_bill_id_foreign` FOREIGN KEY (`bill_id`) REFERENCES `bills` (`id`),
  CONSTRAINT `sales_inventory_id_foreign` FOREIGN KEY (`inventory_id`) REFERENCES `inventories` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sales`
--

LOCK TABLES `sales` WRITE;
/*!40000 ALTER TABLE `sales` DISABLE KEYS */;
/*!40000 ALTER TABLE `sales` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `scales`
--

DROP TABLE IF EXISTS `scales`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `scales` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `description` varchar(255) DEFAULT NULL,
  `unit` varchar(10) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `store_id` int(10) unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `scales_store_id_foreign` (`store_id`),
  CONSTRAINT `scales_store_id_foreign` FOREIGN KEY (`store_id`) REFERENCES `stores` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `scales`
--

LOCK TABLES `scales` WRITE;
/*!40000 ALTER TABLE `scales` DISABLE KEYS */;
/*!40000 ALTER TABLE `scales` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `stock_flows`
--

DROP TABLE IF EXISTS `stock_flows`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `stock_flows` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `flow` double(4,2) NOT NULL DEFAULT 1.00,
  `stockable_id` varchar(255) DEFAULT NULL,
  `stockable_type` varchar(255) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `inventory_id` int(10) unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `stock_flows_inventory_id_foreign` (`inventory_id`),
  CONSTRAINT `stock_flows_inventory_id_foreign` FOREIGN KEY (`inventory_id`) REFERENCES `inventories` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `stock_flows`
--

LOCK TABLES `stock_flows` WRITE;
/*!40000 ALTER TABLE `stock_flows` DISABLE KEYS */;
/*!40000 ALTER TABLE `stock_flows` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `stocks`
--

DROP TABLE IF EXISTS `stocks`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `stocks` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `available_stock` double(8,2) NOT NULL DEFAULT 1.00,
  `inventory_id` int(10) unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `stocks_inventory_id_foreign` (`inventory_id`),
  CONSTRAINT `stocks_inventory_id_foreign` FOREIGN KEY (`inventory_id`) REFERENCES `inventories` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `stocks`
--

LOCK TABLES `stocks` WRITE;
/*!40000 ALTER TABLE `stocks` DISABLE KEYS */;
/*!40000 ALTER TABLE `stocks` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `store_user`
--

DROP TABLE IF EXISTS `store_user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `store_user` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `stores_id` int(10) unsigned DEFAULT NULL,
  `users_id` int(10) unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `store_user` (`stores_id`,`users_id`),
  KEY `store_user_stores_id_foreign` (`stores_id`),
  KEY `store_user_users_id_foreign` (`users_id`),
  CONSTRAINT `store_user_stores_id_foreign` FOREIGN KEY (`stores_id`) REFERENCES `stores` (`id`),
  CONSTRAINT `store_user_users_id_foreign` FOREIGN KEY (`users_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `store_user`
--

LOCK TABLES `store_user` WRITE;
/*!40000 ALTER TABLE `store_user` DISABLE KEYS */;
/*!40000 ALTER TABLE `store_user` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `stores`
--

DROP TABLE IF EXISTS `stores`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `stores` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `store_name` varchar(255) NOT NULL,
  `address` text DEFAULT NULL,
  `contact` tinytext DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name_unique` (`store_name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `stores`
--

LOCK TABLES `stores` WRITE;
/*!40000 ALTER TABLE `stores` DISABLE KEYS */;
/*!40000 ALTER TABLE `stores` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `terminals`
--

DROP TABLE IF EXISTS `terminals`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `terminals` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `machine` varchar(255) NOT NULL,
  `printer` varchar(255) NOT NULL,
  `balance` double(24,8) NOT NULL DEFAULT 1.00000000,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `user_id` int(10) unsigned DEFAULT NULL,
  `store_id` int(10) unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `terminals_user_id_foreign` (`user_id`),
  KEY `terminals_store_id_foreign` (`store_id`),
  CONSTRAINT `terminals_store_id_foreign` FOREIGN KEY (`store_id`) REFERENCES `stores` (`id`),
  CONSTRAINT `terminals_user_id_foreign` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `terminals`
--

LOCK TABLES `terminals` WRITE;
/*!40000 ALTER TABLE `terminals` DISABLE KEYS */;
/*!40000 ALTER TABLE `terminals` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `tokens`
--

DROP TABLE IF EXISTS `tokens`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `tokens` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `key` mediumtext NOT NULL,
  `valid_thru` datetime NOT NULL,
  `valid` tinyint(1) NOT NULL,
  `user_id` int(10) unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `tokens_user_id_foreign` (`user_id`),
  CONSTRAINT `tokens_user_id_foreign` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tokens`
--

LOCK TABLES `tokens` WRITE;
/*!40000 ALTER TABLE `tokens` DISABLE KEYS */;
INSERT INTO `tokens` VALUES (1,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjEsImV4cCI6MTUwMjA3NTA1OSwibmJmIjoxNTAyMDc0NDU5fQ.p_ug83LL6R_0wuc0oT6jsMXULJQfUYj1QIjGPB5QD4E','2017-08-07 08:34:19',1,1);
/*!40000 ALTER TABLE `tokens` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `users` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(255) DEFAULT NULL,
  `email` varchar(255) DEFAULT NULL,
  `password` varchar(255) DEFAULT NULL,
  `remember_token` varchar(255) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `username` (`username`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES (1,'gnanakeethan','gnanakeethan@gmail.com','$2y$10$uZh.v9W3.1t2QJK.Ki6naeh.zG5q/rZZ4JOWRWKLRWpV9qyrl2c62',NULL,'2017-06-02 23:01:19','2017-06-02 23:01:19'),(2,'dasa','admin@pos.run','$2a$15$QIz8Ax1eBH0GQT.tqFmGJuwAedsl33ihwFNgg05PspUND0TOZdArq','',NULL,NULL);
/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2017-08-07  8:26:26
