-- Created with golang-migrates
-- command: migrate create -ext sql -dir [directory] -seq [name]

drop table if exists entries;
drop table if exists transfers;
drop table if exists accounts;