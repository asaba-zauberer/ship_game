SHOW TABLES;
select * from user;
SELECT * from user_collection_item;
SELECT * from gacha_probability;
SELECT * FROM collection_item NATURAL LEFT JOIN (SELECT user_id, collection_item_id AS id FROM user_collection_item WHERE user_id = 'ffe1aa17-357b-4235-829f-9fa730ec5542') AS user_collection;
show columns from collection_item;

delete from user_collection_item where user_id = '5b87cccb-2caf-4947-8617-abf70fa823ec';

select c.id, c.name, c.rarity, u.user_id from collection_item as c left join user_collection_item as u on c.id = u.collection_item_id where id in (1001, 1038	); 