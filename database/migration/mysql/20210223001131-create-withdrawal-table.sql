-- +migrate Up
CREATE TABLE IF NOT EXISTS `flip`.`withdrawls`
(
    `id`
    INT
    UNSIGNED
    NOT
    NULL
    AUTO_INCREMENT,
    `transaction_code`
    INT
    NOT
    NULL,
    `bank_code`
    VARCHAR
(
    45
) NOT NULL,
    `account_number` VARCHAR
(
    45
) NOT NULL,
    `amount` FLOAT NOT NULL,
    `remark` TEXT NOT NULL,
    `status` VARCHAR
(
    45
) NOT NULL,
    `created_at` TIMESTAMP NOT NULL,
    `updated_at` TIMESTAMP NOT NULL,
    PRIMARY KEY
(
    `id`
),
    UNIQUE INDEX `transaction_code_UNIQUE`
(
    `transaction_code` ASC
))
    ENGINE = InnoDB;
-- +migrate Down
DROP TABLE IF EXISTS `flip`.`withdrawls`;