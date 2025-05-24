
-- Table: Customers
CREATE TABLE Customers (
    ID INT AUTO_INCREMENT PRIMARY KEY,
    CustomerID VARCHAR(50) UNIQUE,
    Name VARCHAR(255),
    Email VARCHAR(255),
    Address TEXT
);

-- Table: Products
CREATE TABLE Products (
    ID INT AUTO_INCREMENT PRIMARY KEY,
    ProductID VARCHAR(50) UNIQUE,
    ProductName VARCHAR(255),
    Category VARCHAR(100)
);

-- Table: Orders
CREATE TABLE Orders (
    ID INT AUTO_INCREMENT PRIMARY KEY,
    OrderID INT UNIQUE,
    CustomerID VARCHAR(50),
    DateOfSale DATE,
    Region VARCHAR(100),
    PaymentMethod VARCHAR(50),
    ShippingCost DECIMAL(10, 2),
    FOREIGN KEY (CustomerID) REFERENCES Customers(CustomerID)
        ON UPDATE CASCADE
        ON DELETE RESTRICT
);

-- Table: OrderDetails
CREATE TABLE OrderDetails (
    ID INT AUTO_INCREMENT PRIMARY KEY,
    OrderID INT,
    ProductID VARCHAR(50),
    QuantitySold INT,
    UnitPrice DECIMAL(10, 2),
    Discount DECIMAL(4, 2),
    FOREIGN KEY (OrderID) REFERENCES Orders(OrderID)
        ON UPDATE CASCADE
        ON DELETE CASCADE,
    FOREIGN KEY (ProductID) REFERENCES Products(ProductID)
        ON UPDATE CASCADE
        ON DELETE RESTRICT
);
