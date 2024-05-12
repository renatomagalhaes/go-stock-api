CREATE TABLE stocks (
    id INT AUTO_INCREMENT,
    sku VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    type INT,
    quantity INT NOT NULL,
    PRIMARY KEY (id)
);