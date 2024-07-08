CREATE TABLE pets (
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    species VARCHAR(255),
    pet_size VARCHAR(255),
    name VARCHAR(255),
    average_male_adult_weight INT UNSIGNED,
    average_female_adult_weight INT UNSIGNED
);
