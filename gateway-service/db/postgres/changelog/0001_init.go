package changelog

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(upInit, downInit)
}

func upInit(tx *sql.Tx) error {
	_, err := tx.Exec(`
		create schema service;

		create table service.user (
			id uuid not null, 
			first_name varchar(100) not null,
			last_name varchar(100) not null,
			password varchar(100) not null,
			email varchar(100) not null primary key,
			phone bigint not null,
			user_type varchar(100) not null
		);

		create table service.payment_information (
			operation_id uuid not null primary key,
			user_id uuid not null,
			amount numeric(10,2) not null,
			card_number numeric(10,2) not null,
			card_holder_name varchar(100) not null,
			cvv numeric(10,2) not null,
			payment_status varchar(100) not null
		);
	`)
	if err != nil {
		return err
	}

	return nil
}

func downInit(tx *sql.Tx) error {
	_, err := tx.Exec(`
		delete table service.user;
		delete table service.payment_information;

		drop schema service;
	`)
	if err != nil {
		return err
	}

	return nil
}
