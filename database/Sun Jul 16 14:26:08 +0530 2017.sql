-- MySQL dump 10.16  Distrib 10.2.6-MariaDB, for osx10.12 (x86_64)
--
-- Host: localhost    Database: posres
-- ------------------------------------------------------
-- Server version	10.2.6-MariaDB

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
  `customer_id` int(10) unsigned DEFAULT NULL,
  `cost` double(8,2) DEFAULT 0.00,
  `gross_total` double(8,2) DEFAULT 0.00,
  `discount` double(8,2) DEFAULT 0.00,
  `net_total` double(8,2) DEFAULT 0.00,
  `balance` double(8,2) DEFAULT 0.00,
  `card_paid` double(8,2) DEFAULT 0.00,
  `cash_paid` double(8,2) DEFAULT 0.00,
  `user_id` int(10) unsigned DEFAULT NULL,
  `terminal_id` int(10) unsigned DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `store_id` int(10) unsigned DEFAULT NULL,
  `closed` tinyint(1) NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`),
  KEY `bills_customer_id_foreign` (`customer_id`),
  KEY `bills_user_id_foreign` (`user_id`),
  KEY `bills_terminal_id_foreign` (`terminal_id`),
  KEY `bills_stores_id_fk` (`store_id`),
  CONSTRAINT `bills_customer_id_foreign` FOREIGN KEY (`customer_id`) REFERENCES `customers` (`id`),
  CONSTRAINT `bills_stores_id_fk` FOREIGN KEY (`store_id`) REFERENCES `stores` (`id`),
  CONSTRAINT `bills_terminal_id_foreign` FOREIGN KEY (`terminal_id`) REFERENCES `terminals` (`id`) ON DELETE SET NULL,
  CONSTRAINT `bills_user_id_foreign` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE SET NULL
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `bills`
--

LOCK TABLES `bills` WRITE;
/*!40000 ALTER TABLE `bills` DISABLE KEYS */;
INSERT INTO `bills` VALUES (6,NULL,668.00,504.00,60.48,443.52,-56.48,300.00,200.00,1,1,NULL,'2017-07-09 17:26:18','2017-07-14 02:15:50',1,1),(7,NULL,268.00,104.00,12.48,91.52,91.52,0.00,0.00,1,1,NULL,'2017-07-14 02:16:02','2017-07-14 02:16:02',1,0),(8,NULL,0.00,0.00,0.00,0.00,0.00,0.00,0.00,1,1,NULL,'2017-07-15 13:53:00','2017-07-15 13:53:00',1,0),(9,NULL,0.00,0.00,0.00,0.00,0.00,0.00,0.00,1,1,NULL,'2017-07-15 13:53:27','2017-07-15 13:53:27',1,0);
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
  `name` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `contact_no` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `address` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
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
  `discount_id` int(10) unsigned DEFAULT NULL,
  `discountable_id` int(10) unsigned DEFAULT NULL,
  `discountable_type` varchar(255) NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `discountables_discount_id_foreign` (`discount_id`),
  CONSTRAINT `discountables_discount_id_foreign` FOREIGN KEY (`discount_id`) REFERENCES `discounts` (`id`) ON DELETE SET NULL
) ENGINE=InnoDB AUTO_INCREMENT=21 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `discountables`
--

LOCK TABLES `discountables` WRITE;
/*!40000 ALTER TABLE `discountables` DISABLE KEYS */;
INSERT INTO `discountables` VALUES (17,1,7,'Products','2017-07-10 10:27:00','2017-07-10 10:27:00'),(18,1,6,'Inventories','2017-07-10 10:35:24','2017-07-10 10:35:24'),(19,1,5,'Inventories','2017-07-10 10:56:27','2017-07-10 10:56:27');
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
  `name` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `from` datetime NOT NULL,
  `to` datetime NOT NULL,
  `type` enum('value','percentage') COLLATE utf8_unicode_ci NOT NULL,
  `value` double(8,2) NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `store_id` int(10) unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `discounts_stores_id_fk` (`store_id`),
  CONSTRAINT `discounts_stores_id_fk` FOREIGN KEY (`store_id`) REFERENCES `stores` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `discounts`
--

LOCK TABLES `discounts` WRITE;
/*!40000 ALTER TABLE `discounts` DISABLE KEYS */;
INSERT INTO `discounts` VALUES (1,'Testing Discot','2017-07-04 00:00:00','2017-07-25 00:00:00','percentage',12.00,'2017-07-09 17:17:27','2017-07-10 11:02:29',1);
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
  `batch_no` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `expiry` date DEFAULT NULL,
  `product_id` int(10) unsigned NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `service` tinyint(1) DEFAULT 0,
  PRIMARY KEY (`id`),
  KEY `inventories_product_id_foreign` (`product_id`),
  CONSTRAINT `inventories_product_id_foreign` FOREIGN KEY (`product_id`) REFERENCES `products` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `inventories`
--

LOCK TABLES `inventories` WRITE;
/*!40000 ALTER TABLE `inventories` DISABLE KEYS */;
INSERT INTO `inventories` VALUES (4,'NANA','0000-12-31',7,'2017-07-09 17:29:52','2017-07-09 17:29:55',0),(5,'Rs.50','0201-07-18',6,'2017-07-09 17:33:00','2017-07-15 02:19:07',0),(6,'Rs.52','0001-02-12',6,'2017-07-10 04:22:28','2017-07-10 04:22:50',0),(7,'POA','0000-12-31',7,'2017-07-10 11:02:49','2017-07-10 11:02:51',0);
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
  `inventory_id` int(10) unsigned NOT NULL,
  `scale_id` int(10) unsigned NOT NULL,
  `units` double(8,2) NOT NULL,
  `price` double(8,2) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `inventory_scale_inventory_id_index` (`inventory_id`),
  KEY `inventory_scale_scale_id_index` (`scale_id`),
  CONSTRAINT `inventory_scale_inventory_id_foreign` FOREIGN KEY (`inventory_id`) REFERENCES `inventories` (`id`) ON DELETE CASCADE,
  CONSTRAINT `inventory_scale_scale_id_foreign` FOREIGN KEY (`scale_id`) REFERENCES `scales` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `inventory_scale`
--

LOCK TABLES `inventory_scale` WRITE;
/*!40000 ALTER TABLE `inventory_scale` DISABLE KEYS */;
INSERT INTO `inventory_scale` VALUES (2,4,1,1.00,200.00),(3,5,1,1.00,50.00),(5,6,1,1.00,52.00),(7,7,1,1.00,180.00);
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
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `migrations`
--

LOCK TABLES `migrations` WRITE;
/*!40000 ALTER TABLE `migrations` DISABLE KEYS */;
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
  `role_id` int(10) unsigned NOT NULL,
  `permission_id` int(10) unsigned NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `permission_role_permission_id_role_id_pk` (`permission_id`,`role_id`),
  KEY `permission_role_role_id_foreign` (`role_id`),
  CONSTRAINT `permission_role_permissions_id_fk` FOREIGN KEY (`permission_id`) REFERENCES `permissions` (`id`),
  CONSTRAINT `permission_role_roles_id_fk` FOREIGN KEY (`role_id`) REFERENCES `bmc`.`roles` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1043 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `permission_role`
--

LOCK TABLES `permission_role` WRITE;
/*!40000 ALTER TABLE `permission_role` DISABLE KEYS */;
INSERT INTO `permission_role` VALUES (1036,2,4),(74,1,12),(1018,2,12),(1038,2,18),(1022,2,21),(1024,2,31),(77,1,47),(1020,2,47),(1028,2,49),(75,1,62),(1026,2,66),(76,1,69),(1030,2,79),(1032,2,80),(1040,2,86),(78,1,90),(1034,2,90),(1016,2,158);
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
  `name` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `display_name` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `description` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `permissions_name_unique` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=412 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `permissions`
--

LOCK TABLES `permissions` WRITE;
/*!40000 ALTER TABLE `permissions` DISABLE KEYS */;
INSERT INTO `permissions` VALUES (1,'inventories_id_get','Inventories Id Get','CAN RETRIEVE inventories/id/','2017-07-15 13:50:45','2017-07-15 13:50:45'),(2,'terminals_id_delete','Terminals Id Delete','CAN DELETE terminals/id/','2017-07-15 13:50:45','2017-07-15 13:50:45'),(3,'jobs_id_delete','Jobs Id Delete','CAN DELETE jobs/id/','2017-07-15 13:50:45','2017-07-15 13:50:45'),(4,'users_id_get','Users Id Get','CAN RETRIEVE users/id/','2017-07-15 13:50:45','2017-07-15 13:50:45'),(5,'users_get','Users Get','CAN RETRIEVE users/','2017-07-15 13:50:45','2017-07-15 13:50:45'),(6,'inventory_scale_get','Inventory Scale Get','CAN RETRIEVE inventory/scale/','2017-07-15 13:50:45','2017-07-15 13:50:45'),(7,'inventory_scale_id_delete','Inventory Scale Id Delete','CAN DELETE inventory/scale/id/','2017-07-15 13:50:45','2017-07-15 13:50:45'),(8,'stock_flows_id_put','Stock Flows Id Put','CAN UPDATE stock/flows/id/','2017-07-15 13:50:45','2017-07-15 13:50:45'),(9,'inventories_get','Inventories Get','CAN RETRIEVE inventories/','2017-07-15 13:50:45','2017-07-15 13:50:45'),(10,'permissions_post','Permissions Post','CAN CREATE permissions/','2017-07-15 13:50:45','2017-07-15 13:50:45'),(11,'discountables_id_delete','Discountables Id Delete','CAN DELETE discountables/id/','2017-07-15 13:50:45','2017-07-15 13:50:45'),(12,'sales_id_get','Sales Id Get','CAN RETRIEVE sales/id/','2017-07-15 13:50:45','2017-07-15 13:50:45'),(13,'terminals_get','Terminals Get','CAN RETRIEVE terminals/','2017-07-15 13:50:45','2017-07-15 13:50:45'),(14,'customers_get','Customers Get','CAN RETRIEVE customers/','2017-07-15 13:50:45','2017-07-15 13:50:45'),(15,'purchases_id_put','Purchases Id Put','CAN UPDATE purchases/id/','2017-07-15 13:50:45','2017-07-15 13:50:45'),(16,'customers_post','Customers Post','CAN CREATE customers/','2017-07-15 13:50:45','2017-07-15 13:50:45'),(17,'inventories_id_delete','Inventories Id Delete','CAN DELETE inventories/id/','2017-07-15 13:50:45','2017-07-15 13:50:45'),(18,'terminals_empty_get','Terminals Empty Get','CAN RETRIEVE terminals/empty/','2017-07-15 13:50:45','2017-07-15 13:50:45'),(19,'purchases_id_get','Purchases Id Get','CAN RETRIEVE purchases/id/','2017-07-15 13:50:45','2017-07-15 13:50:45'),(20,'_routes_get',' Routes Get','CAN RETRIEVE /routes/','2017-07-15 13:50:45','2017-07-15 13:50:45'),(21,'bills_get','Bills Get','CAN RETRIEVE bills/','2017-07-15 13:50:45','2017-07-15 13:50:45'),(22,'stocks_get','Stocks Get','CAN RETRIEVE stocks/','2017-07-15 13:50:45','2017-07-15 13:50:45'),(23,'discountables_id_get','Discountables Id Get','CAN RETRIEVE discountables/id/','2017-07-15 13:50:45','2017-07-15 13:50:45'),(24,'products_id_put','Products Id Put','CAN UPDATE products/id/','2017-07-15 13:50:45','2017-07-15 13:50:45'),(25,'inventory_scale_id_put','Inventory Scale Id Put','CAN UPDATE inventory/scale/id/','2017-07-15 13:50:45','2017-07-15 13:50:45'),(26,'stocks_post','Stocks Post','CAN CREATE stocks/','2017-07-15 13:50:45','2017-07-15 13:50:45'),(27,'customers_id_delete','Customers Id Delete','CAN DELETE customers/id/','2017-07-15 13:50:45','2017-07-15 13:50:45'),(28,'discounts_id_delete','Discounts Id Delete','CAN DELETE discounts/id/','2017-07-15 13:50:45','2017-07-15 13:50:45'),(29,'discounts_id_get','Discounts Id Get','CAN RETRIEVE discounts/id/','2017-07-15 13:50:45','2017-07-15 13:50:45'),(30,'jobs_id_get','Jobs Id Get','CAN RETRIEVE jobs/id/','2017-07-15 13:50:45','2017-07-15 13:50:45'),(31,'bills_id_put','Bills Id Put','CAN UPDATE bills/id/','2017-07-15 13:50:45','2017-07-15 13:50:45'),(32,'terminals_post','Terminals Post','CAN CREATE terminals/','2017-07-15 13:50:45','2017-07-15 13:50:45'),(33,'auth_logout_post','Auth Logout Post','CAN CREATE auth/logout/','2017-07-15 13:50:45','2017-07-15 13:50:45'),(34,'stores_post','Stores Post','CAN CREATE stores/','2017-07-15 13:50:45','2017-07-15 13:50:45'),(35,'products_id_delete','Products Id Delete','CAN DELETE products/id/','2017-07-15 13:50:45','2017-07-15 13:50:45'),(36,'_get',' Get','CAN RETRIEVE /','2017-07-15 13:50:45','2017-07-15 13:50:45'),(37,'discounts_get','Discounts Get','CAN RETRIEVE discounts/','2017-07-15 13:50:45','2017-07-15 13:50:45'),(38,'stock_flows_get','Stock Flows Get','CAN RETRIEVE stock/flows/','2017-07-15 13:50:45','2017-07-15 13:50:45'),(39,'permissions_update_permissions_get','Permissions Update Permissions Get','CAN RETRIEVE permissions/update/permissions/','2017-07-15 13:50:45','2017-07-15 13:50:45'),(40,'inventories_post','Inventories Post','CAN CREATE inventories/','2017-07-15 13:50:45','2017-07-15 13:50:45'),(41,'users_post','Users Post','CAN CREATE users/','2017-07-15 13:50:45','2017-07-15 13:50:45'),(42,'stores_id_get','Stores Id Get','CAN RETRIEVE stores/id/','2017-07-15 13:50:45','2017-07-15 13:50:45'),(43,'stores_get','Stores Get','CAN RETRIEVE stores/','2017-07-15 13:50:45','2017-07-15 13:50:45'),(44,'inventory_scale_id_get','Inventory Scale Id Get','CAN RETRIEVE inventory/scale/id/','2017-07-15 13:50:45','2017-07-15 13:50:45'),(45,'users_id_put','Users Id Put','CAN UPDATE users/id/','2017-07-15 13:50:45','2017-07-15 13:50:45'),(46,'scales_id_delete','Scales Id Delete','CAN DELETE scales/id/','2017-07-15 13:50:45','2017-07-15 13:50:45'),(47,'sales_get','Sales Get','CAN RETRIEVE sales/','2017-07-15 13:50:45','2017-07-15 13:50:45'),(48,'roles_id_get','Roles Id Get','CAN RETRIEVE roles/id/','2017-07-15 13:50:45','2017-07-15 13:50:45'),(49,'products_bill_get','Products Bill Get','CAN RETRIEVE products/bill/','2017-07-15 13:50:45','2017-07-15 13:50:45'),(50,'stocks_id_put','Stocks Id Put','CAN UPDATE stocks/id/','2017-07-15 13:50:45','2017-07-15 13:50:45'),(51,'discountables_post','Discountables Post','CAN CREATE discountables/','2017-07-15 13:50:45','2017-07-15 13:50:45'),(52,'roles_id_delete','Roles Id Delete','CAN DELETE roles/id/','2017-07-15 13:50:45','2017-07-15 13:50:45'),(53,'stocks_id_get','Stocks Id Get','CAN RETRIEVE stocks/id/','2017-07-15 13:50:45','2017-07-15 13:50:45'),(54,'stock_flows_id_get','Stock Flows Id Get','CAN RETRIEVE stock/flows/id/','2017-07-15 13:50:45','2017-07-15 13:50:45'),(55,'permissions_get','Permissions Get','CAN RETRIEVE permissions/','2017-07-15 13:50:45','2017-07-15 13:50:45'),(56,'products_get','Products Get','CAN RETRIEVE products/','2017-07-15 13:50:45','2017-07-15 13:50:45'),(57,'products_id_get','Products Id Get','CAN RETRIEVE products/id/','2017-07-15 13:50:45','2017-07-15 13:50:45'),(58,'terminals_id_put','Terminals Id Put','CAN UPDATE terminals/id/','2017-07-15 13:50:45','2017-07-15 13:50:45'),(59,'discounts_id_put','Discounts Id Put','CAN UPDATE discounts/id/','2017-07-15 13:50:45','2017-07-15 13:50:45'),(60,'discounts_post','Discounts Post','CAN CREATE discounts/','2017-07-15 13:50:45','2017-07-15 13:50:45'),(61,'customers_id_get','Customers Id Get','CAN RETRIEVE customers/id/','2017-07-15 13:50:45','2017-07-15 13:50:45'),(62,'sales_id_put','Sales Id Put','CAN UPDATE sales/id/','2017-07-15 13:50:45','2017-07-15 13:50:45'),(63,'scales_id_put','Scales Id Put','CAN UPDATE scales/id/','2017-07-15 13:50:45','2017-07-15 13:50:45'),(64,'roles_id_put','Roles Id Put','CAN UPDATE roles/id/','2017-07-15 13:50:45','2017-07-15 13:50:45'),(65,'inventory_scale_post','Inventory Scale Post','CAN CREATE inventory/scale/','2017-07-15 13:50:45','2017-07-15 13:50:45'),(66,'bills_post','Bills Post','CAN CREATE bills/','2017-07-15 13:50:45','2017-07-15 13:50:45'),(67,'users_id_delete','Users Id Delete','CAN DELETE users/id/','2017-07-15 13:50:45','2017-07-15 13:50:45'),(68,'stores_id_delete','Stores Id Delete','CAN DELETE stores/id/','2017-07-15 13:50:45','2017-07-15 13:50:45'),(69,'sales_id_delete','Sales Id Delete','CAN DELETE sales/id/','2017-07-15 13:50:45','2017-07-15 13:50:45'),(70,'printer_id_get','Printer Id Get','CAN RETRIEVE printer/id/','2017-07-15 13:50:45','2017-07-15 13:50:45'),(71,'roles_get','Roles Get','CAN RETRIEVE roles/','2017-07-15 13:50:45','2017-07-15 13:50:45'),(72,'discountables_id_put','Discountables Id Put','CAN UPDATE discountables/id/','2017-07-15 13:50:45','2017-07-15 13:50:45'),(73,'customers_id_put','Customers Id Put','CAN UPDATE customers/id/','2017-07-15 13:50:45','2017-07-15 13:50:45'),(74,'inventories_id_put','Inventories Id Put','CAN UPDATE inventories/id/','2017-07-15 13:50:45','2017-07-15 13:50:45'),(75,'auth_validate_post','Auth Validate Post','CAN CREATE auth/validate/','2017-07-15 13:50:45','2017-07-15 13:50:45'),(76,'purchases_id_delete','Purchases Id Delete','CAN DELETE purchases/id/','2017-07-15 13:50:45','2017-07-15 13:50:45'),(77,'permissions_id_delete','Permissions Id Delete','CAN DELETE permissions/id/','2017-07-15 13:50:45','2017-07-15 13:50:45'),(78,'scales_id_get','Scales Id Get','CAN RETRIEVE scales/id/','2017-07-15 13:50:45','2017-07-15 13:50:45'),(79,'bills_payable_get','Bills Payable Get','CAN RETRIEVE bills/payable/','2017-07-15 13:50:45','2017-07-15 13:50:45'),(80,'bills_id_get','Bills Id Get','CAN RETRIEVE bills/id/','2017-07-15 13:50:45','2017-07-15 13:50:45'),(81,'stores_id_put','Stores Id Put','CAN UPDATE stores/id/','2017-07-15 13:50:45','2017-07-15 13:50:45'),(82,'roles_post','Roles Post','CAN CREATE roles/','2017-07-15 13:50:45','2017-07-15 13:50:45'),(83,'products_post','Products Post','CAN CREATE products/','2017-07-15 13:50:45','2017-07-15 13:50:45'),(84,'bills_id_delete','Bills Id Delete','CAN DELETE bills/id/','2017-07-15 13:50:45','2017-07-15 13:50:45'),(85,'users_id_permissions_get','Users Id Permissions Get','CAN RETRIEVE users/id/permissions/','2017-07-15 13:50:45','2017-07-15 13:50:45'),(86,'scales_get','Scales Get','CAN RETRIEVE scales/','2017-07-15 13:50:45','2017-07-15 13:50:45'),(87,'jobs_id_put','Jobs Id Put','CAN UPDATE jobs/id/','2017-07-15 13:50:45','2017-07-15 13:50:45'),(88,'jobs_post','Jobs Post','CAN CREATE jobs/','2017-07-15 13:50:45','2017-07-15 13:50:45'),(89,'stock_flows_post','Stock Flows Post','CAN CREATE stock/flows/','2017-07-15 13:50:45','2017-07-15 13:50:45'),(90,'sales_post','Sales Post','CAN CREATE sales/','2017-07-15 13:50:45','2017-07-15 13:50:45'),(91,'auth_post','Auth Post','CAN CREATE auth/','2017-07-15 13:50:45','2017-07-15 13:50:45'),(92,'stock_flows_id_delete','Stock Flows Id Delete','CAN DELETE stock/flows/id/','2017-07-15 13:50:45','2017-07-15 13:50:45'),(93,'terminals_id_get','Terminals Id Get','CAN RETRIEVE terminals/id/','2017-07-15 13:50:45','2017-07-15 13:50:45'),(94,'permissions_id_get','Permissions Id Get','CAN RETRIEVE permissions/id/','2017-07-15 13:50:45','2017-07-15 13:50:45'),(95,'purchases_get','Purchases Get','CAN RETRIEVE purchases/','2017-07-15 13:50:45','2017-07-15 13:50:45'),(96,'jobs_get','Jobs Get','CAN RETRIEVE jobs/','2017-07-15 13:50:45','2017-07-15 13:50:45'),(97,'discountables_get','Discountables Get','CAN RETRIEVE discountables/','2017-07-15 13:50:45','2017-07-15 13:50:45'),(98,'permissions_id_put','Permissions Id Put','CAN UPDATE permissions/id/','2017-07-15 13:50:45','2017-07-15 13:50:45'),(99,'scales_post','Scales Post','CAN CREATE scales/','2017-07-15 13:50:45','2017-07-15 13:50:45'),(100,'auth_refresh_post','Auth Refresh Post','CAN CREATE auth/refresh/','2017-07-15 13:50:45','2017-07-15 13:50:45'),(101,'purchases_post','Purchases Post','CAN CREATE purchases/','2017-07-15 13:50:45','2017-07-15 13:50:45'),(102,'stocks_id_delete','Stocks Id Delete','CAN DELETE stocks/id/','2017-07-15 13:50:45','2017-07-15 13:50:45'),(158,'terminals_request_id_get','Terminals Request Id Get','CAN RETRIEVE terminals/request/id/','2017-07-15 13:50:45','2017-07-15 13:50:45');
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
  `parent_id` int(10) unsigned NOT NULL,
  `child_id` int(10) unsigned NOT NULL,
  `units` double(8,2) DEFAULT 1.00,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `product_product_parent_id_index` (`parent_id`),
  KEY `product_product_child_id_index` (`child_id`),
  CONSTRAINT `product_product_child_id_foreign` FOREIGN KEY (`child_id`) REFERENCES `products` (`id`) ON DELETE CASCADE,
  CONSTRAINT `product_product_parent_id_foreign` FOREIGN KEY (`parent_id`) REFERENCES `products` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
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
  `product_code` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `barcode` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `name` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `singular` tinyint(1) NOT NULL,
  `priority` tinyint(1) DEFAULT NULL,
  `ordering` int(10) unsigned DEFAULT NULL,
  `image` blob DEFAULT NULL,
  `image_type` text COLLATE utf8_unicode_ci DEFAULT NULL,
  `scale_id` int(10) unsigned NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `service` tinyint(1) DEFAULT 0,
  `store_id` int(10) unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `products_scale_id_foreign` (`scale_id`),
  KEY `products_stores_id_fk` (`store_id`),
  CONSTRAINT `products_scale_id_foreign` FOREIGN KEY (`scale_id`) REFERENCES `scales` (`id`) ON DELETE CASCADE,
  CONSTRAINT `products_stores_id_fk` FOREIGN KEY (`store_id`) REFERENCES `stores` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `products`
--

LOCK TABLES `products` WRITE;
/*!40000 ALTER TABLE `products` DISABLE KEYS */;
INSERT INTO `products` VALUES (6,'P213','12312333123','Samaposha',1,0,0,'','',1,NULL,NULL,0,1),(7,'P123','2178312546789','Classisc',1,0,0,'','',1,NULL,NULL,0,1);
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
  `inventory_id` int(10) unsigned DEFAULT NULL,
  `average_cost` double(12,4) NOT NULL,
  `units` double(12,4) NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `stock_flow_id` int(10) unsigned DEFAULT NULL,
  `store_id` int(10) unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `purchases_inventory_id_foreign` (`inventory_id`),
  KEY `purchases_stock_flows_id_fk` (`stock_flow_id`),
  KEY `purchases_stores_id_fk` (`store_id`),
  CONSTRAINT `purchases_inventory_id_foreign` FOREIGN KEY (`inventory_id`) REFERENCES `inventories` (`id`) ON DELETE SET NULL,
  CONSTRAINT `purchases_stock_flows_id_fk` FOREIGN KEY (`stock_flow_id`) REFERENCES `stock_flows` (`id`),
  CONSTRAINT `purchases_stores_id_fk` FOREIGN KEY (`store_id`) REFERENCES `stores` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `purchases`
--

LOCK TABLES `purchases` WRITE;
/*!40000 ALTER TABLE `purchases` DISABLE KEYS */;
INSERT INTO `purchases` VALUES (1,5,180.0000,200.0000,NULL,NULL,NULL,1),(2,4,200.0000,200.0000,NULL,NULL,NULL,1),(3,7,120.0000,200.0000,NULL,NULL,NULL,1),(4,6,134.0000,20.0000,NULL,NULL,NULL,1);
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
  `role_id` int(10) unsigned NOT NULL,
  `user_id` int(10) unsigned NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `role_user_user_id_role_id_pk` (`user_id`,`role_id`),
  KEY `role_user_roles_id_fk` (`role_id`),
  CONSTRAINT `role_user_roles_id_fk` FOREIGN KEY (`role_id`) REFERENCES `bmc`.`roles` (`id`),
  CONSTRAINT `role_user_users_id_fk` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=54 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `role_user`
--

LOCK TABLES `role_user` WRITE;
/*!40000 ALTER TABLE `role_user` DISABLE KEYS */;
INSERT INTO `role_user` VALUES (1,1,1),(53,2,2);
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
  `name` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `display_name` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `description` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `roles_name_unique` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `roles`
--

LOCK TABLES `roles` WRITE;
/*!40000 ALTER TABLE `roles` DISABLE KEYS */;
INSERT INTO `roles` VALUES (1,'superadmin','SuperAdmin','',NULL,NULL),(2,'Master','Blaster','123',NULL,NULL);
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
  `bill_id` int(10) unsigned DEFAULT NULL,
  `inventory_id` int(10) unsigned DEFAULT NULL,
  `units` double(8,4) NOT NULL,
  `cost` double(8,2) DEFAULT 0.00,
  `hide` tinyint(1) DEFAULT 0,
  `unit_price` double(8,2) NOT NULL,
  `total` double(8,2) NOT NULL,
  `discount` double(8,2) NOT NULL,
  `amount` double(8,2) NOT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `sales_bill_id_foreign` (`bill_id`),
  KEY `sales_inventory_id_foreign` (`inventory_id`),
  CONSTRAINT `sales_bill_id_foreign` FOREIGN KEY (`bill_id`) REFERENCES `bills` (`id`) ON DELETE SET NULL,
  CONSTRAINT `sales_inventory_id_foreign` FOREIGN KEY (`inventory_id`) REFERENCES `inventories` (`id`) ON DELETE SET NULL
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sales`
--

LOCK TABLES `sales` WRITE;
/*!40000 ALTER TABLE `sales` DISABLE KEYS */;
INSERT INTO `sales` VALUES (10,6,4,2.0000,400.00,0,200.00,400.00,48.00,352.00,NULL,NULL,NULL),(12,6,6,2.0000,268.00,0,52.00,104.00,12.48,91.52,NULL,NULL,NULL),(13,7,6,2.0000,268.00,0,52.00,104.00,12.48,91.52,NULL,NULL,NULL);
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
  `unit` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `description` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `store_id` int(10) unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `scales_stores_id_fk` (`store_id`),
  CONSTRAINT `scales_stores_id_fk` FOREIGN KEY (`store_id`) REFERENCES `stores` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `scales`
--

LOCK TABLES `scales` WRITE;
/*!40000 ALTER TABLE `scales` DISABLE KEYS */;
INSERT INTO `scales` VALUES (1,'Pcs','Piece',NULL,NULL,1);
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
  `flow` double(8,2) NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `stock_id` int(10) unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `stock_flows_stocks_id_fk` (`stock_id`),
  CONSTRAINT `stock_flows_stocks_id_fk` FOREIGN KEY (`stock_id`) REFERENCES `stocks` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
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
  `inventory_id` int(10) unsigned DEFAULT NULL,
  `available_stock` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `stocks_inventories_id_fk` (`inventory_id`),
  CONSTRAINT `stocks_inventories_id_fk` FOREIGN KEY (`inventory_id`) REFERENCES `inventories` (`id`) ON DELETE CASCADE
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
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `stores_id` int(10) unsigned NOT NULL,
  `users_id` int(10) unsigned NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `store_user_store_id_user_id_uindex` (`stores_id`,`users_id`),
  KEY `store_user_users_id_fk` (`users_id`),
  CONSTRAINT `store_user_stores_id_fk` FOREIGN KEY (`stores_id`) REFERENCES `stores` (`id`),
  CONSTRAINT `store_user_users_id_fk` FOREIGN KEY (`users_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=71 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `store_user`
--

LOCK TABLES `store_user` WRITE;
/*!40000 ALTER TABLE `store_user` DISABLE KEYS */;
INSERT INTO `store_user` VALUES (1,1,1),(70,1,2);
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
  `store_name` varchar(255) DEFAULT NULL,
  `address` text DEFAULT NULL,
  `contact_no` tinytext DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `stores`
--

LOCK TABLES `stores` WRITE;
/*!40000 ALTER TABLE `stores` DISABLE KEYS */;
INSERT INTO `stores` VALUES (1,'GTech','Kondavil','215367890');
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
  `user_id` int(10) unsigned DEFAULT NULL,
  `machine` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `printer` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `balance` double(24,2) NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `store_id` int(10) unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `terminals_user_id_foreign` (`user_id`),
  KEY `terminals_stores_id_fk` (`store_id`),
  CONSTRAINT `terminals_stores_id_fk` FOREIGN KEY (`store_id`) REFERENCES `stores` (`id`),
  CONSTRAINT `terminals_user_id_foreign` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE SET NULL
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `terminals`
--

LOCK TABLES `terminals` WRITE;
/*!40000 ALTER TABLE `terminals` DISABLE KEYS */;
INSERT INTO `terminals` VALUES (1,NULL,'POS','OSP',0.00,NULL,NULL,1),(2,NULL,'PSD','DSA',0.00,NULL,NULL,1);
/*!40000 ALTER TABLE `terminals` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `tokens`
--

DROP TABLE IF EXISTS `tokens`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `tokens` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `key` varchar(255) DEFAULT NULL,
  `valid_thru` datetime DEFAULT NULL,
  `valid` tinyint(1) DEFAULT 1,
  `user_id` int(10) unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `invalid_tokens_id_uindex` (`id`),
  KEY `tokens_users_id_fk` (`user_id`),
  CONSTRAINT `tokens_users_id_fk` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1366 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tokens`
--

LOCK TABLES `tokens` WRITE;
/*!40000 ALTER TABLE `tokens` DISABLE KEYS */;
INSERT INTO `tokens` VALUES (1364,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjIsImV4cCI6MTUwMDE3ODQ5MywibmJmIjoxNTAwMTc3ODkzfQ.tgOhsCprBUCFgEwAmsPY3FrXkClN4SIJlVph8RI6Olo','2017-07-16 09:44:53',0,2),(1365,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjIsImV4cCI6MTUwMDE3ODgxNCwibmJmIjoxNTAwMTc4MjE0fQ.wTmfcrFt3eTF2iCXh9c4WoYnTx9_N78_sivdUQBNifQ','2017-07-16 09:50:14',1,2);
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
  `username` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `email` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `password` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `remember_token` varchar(100) COLLATE utf8_unicode_ci DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `users_email_unique` (`email`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
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

-- Dump completed on 2017-07-16 14:26:08
