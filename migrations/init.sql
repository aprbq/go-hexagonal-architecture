# sql script

-- สร้างตาราง Customers
CREATE TABLE Customers (
    customer_id INT PRIMARY KEY,
    name VARCHAR(100),
    date_of_birth DATE,
    city VARCHAR(50),
    zipcode VARCHAR(20),
    status INT
);

-- เพิ่มข้อมูลตาราง Customers
INSERT INTO Customers (customer_id, name, date_of_birth, city, zipcode, status) 
VALUES
    (2000, 'Steve', '1978-12-15', 'Delhi', '110075', 1),
    (2001, 'Arian', '1988-05-21', 'Newburgh, NY', '12550', 1),
    (2002, 'Hadley', '1988-04-30', 'Englewood, NJ', '07631', 1),
    (2003, 'Ben', '1988-01-04', 'Manchester, NH', '03102', 0),
    (2004, 'Nina', '1988-05-14', 'Clarkston, MI', '48348', 1),
    (2005, 'Osman', '1988-11-08', 'Hyattsville, MD', '20782', 0);

-- สร้างตาราง Accounts
CREATE TABLE Accounts (
    account_id INT PRIMARY KEY,
    customer_id INT,
    opening_date TIMESTAMP,
    account_type VARCHAR(20),
    amount DECIMAL(15, 2),
    status INT
);

-- เพิ่มข้อมูลตาราง Accounts
INSERT INTO Accounts (account_id, customer_id, opening_date, account_type, amount, status) 
VALUES
    (95470, 2000, '2020-08-22 10:20:06', 'saving', 6823.23, 1),
    (95471, 2002, '2020-08-09 10:27:22', 'checking', 3342.96, 1),
    (95472, 2001, '2020-08-09 10:35:22', 'saving', 7000.00, 1),
    (95473, 2001, '2020-08-09 10:38:22', 'saving', 5861.86, 1);