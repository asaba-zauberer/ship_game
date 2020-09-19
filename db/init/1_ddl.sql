SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';

-- -----------------------------------------------------
-- Schema ca_hack
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `ca_hack` DEFAULT CHARACTER SET utf8mb4 ;
USE `ca_hack` ;

SET CHARSET utf8mb4;

-- -----------------------------------------------------
-- Table `ca_hack`.`user`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `ca_hack`.`user` (
  `id` VARCHAR(128) NOT NULL COMMENT 'ユーザID',
  `auth_token` VARCHAR(128) NOT NULL COMMENT '認証トークン',
  `name` VARCHAR(64) NOT NULL COMMENT 'ユーザ名',
  `coin` INT UNSIGNED NOT NULL COMMENT '所持コイン',
  `stage` INT UNSIGNED NOT NULL COMMENT '到達ステージ',
  PRIMARY KEY (`id`),
  INDEX `idx_auth_token` (`auth_token` ASC))
ENGINE = InnoDB
COMMENT = 'ユーザ';


-- -----------------------------------------------------
-- Table `ca_hack`.`collection_item`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `ca_hack`.`collection_item` (
  `id` VARCHAR(128) NOT NULL COMMENT 'コレクションID',
  `name` VARCHAR(64) NOT NULL COMMENT 'コレクション名',
  PRIMARY KEY (`id`))
ENGINE = InnoDB
COMMENT = 'コレクションアイテム';

-- -----------------------------------------------------
-- Table `ca_hack`.`user_score`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `ca_hack`.`user_score` (
  `id` VARCHAR(128) NOT NULL COMMENT 'ユーザID',
  `stage` INT UNSIGNED NOT NULL COMMENT 'ステージ',
  `score` INT UNSIGNED NOT NULL COMMENT 'スコア',
  PRIMARY KEY (`id`),
  INDEX `idx_score` (`score` ASC),
  CONSTRAINT `fk_id_user`
    FOREIGN KEY (`id`)
    REFERENCES `ca_hack`.`user` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB
COMMENT = 'ハイスコア';


-- -----------------------------------------------------
-- Table `ca_hack`.`user_collection_item`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `ca_hack`.`user_collection_item` (
  `user_id` VARCHAR(128) NOT NULL COMMENT 'ユーザID',
  `collection_item_id` VARCHAR(128) NOT NULL COMMENT 'コレクションアイテムID',
  PRIMARY KEY (`user_id`, `collection_item_id`),
  INDEX `fk_user_collection_item_user_idx` (`user_id` ASC),
  INDEX `fk_user_collection_item_collection_item_idx` (`collection_item_id` ASC),
  CONSTRAINT `fk_user_collection_item_user`
    FOREIGN KEY (`user_id`)
    REFERENCES `ca_hack`.`user` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_user_collection_item_collection_item`
    FOREIGN KEY (`collection_item_id`)
    REFERENCES `ca_hack`.`collection_item` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB
COMMENT = 'ユーザ所持コレクションアイテム';


-- -----------------------------------------------------
-- Table `ca_hack`.`gacha_probability`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `ca_hack`.`gacha_probability` (
  `collection_item_id` VARCHAR(128) NOT NULL COMMENT 'コレクションアイテムID',
  `ratio` INT UNSIGNED NOT NULL COMMENT '排出重み',
  INDEX `fk_gacha_probability_collection_item_idx` (`collection_item_id` ASC),
  PRIMARY KEY (`collection_item_id`),
  CONSTRAINT `fk_gacha_probability_collection_item_id`
    FOREIGN KEY (`collection_item_id`)
    REFERENCES `ca_hack`.`collection_item` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB
COMMENT = 'ガチャ排出情報';


SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;
