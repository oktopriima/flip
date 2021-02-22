-- +migrate Up
CREATE TABLE IF NOT EXISTS `flip`.`withdrawals`
(
    `id`               INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `transaction_code` BIGINT(255)  NOT NULL,
    `bank_code`        VARCHAR(45)  NOT NULL,
    `account_number`   VARCHAR(45)  NOT NULL,
    `amount`           FLOAT        NOT NULL,
    `remark`           TEXT         NOT NULL,
    `status`           VARCHAR(45)  NOT NULL,
    `receipt`          TEXT         NOT NULL,
    `time_served`      VARCHAR(45)  NOT NULL,
    `created_at`       TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`       TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE INDEX `transaction_code_UNIQUE` (`transaction_code` ASC)
)
    ENGINE = InnoDB;
-- +migrate Down
DROP TABLE IF EXISTS `flip`.`withdrawals`;