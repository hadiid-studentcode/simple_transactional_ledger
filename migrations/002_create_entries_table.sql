CREATE TABLE entries (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    account_id BIGINT,
    amount DECIMAL(10,2) NOT NULL,
    create_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    update_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (account_id) REFERENCES accounts(id) ON DELETE CASCADE ON UPDATE CASCADE
    
)

SELECT * FROM accounts

-- INSERT
INSERT INTO entries (account_id, amount) VALUES (12, 100.00);

-- UPDATE
UPDATE entries SET amount = 200.00 WHERE id = 1;

-- DELETE
DELETE FROM entries WHERE id = 2;

-- SELECT
SELECT * FROM entries

-- select where id = 3 join with accounts table
SELECT e.id, e.account_id, e.amount, e.create_at, e.update_at, a.name, a.balance
FROM entries e
JOIN accounts a ON e.account_id = a.id
WHERE e.id = 3


SELECT e.id, e.account_id, e.amount, e.create_at, e.update_at, a.name, a.balance
		FROM entries e
		JOIN accounts a ON e.account_id = a.id

