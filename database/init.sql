-- Create tables
CREATE TABLE Airports (
    airport_id SERIAL PRIMARY KEY,
    name VARCHAR(100),
    city VARCHAR(50),
    IATA_code VARCHAR(10),
    ICAO_code VARCHAR(10)
);

CREATE TABLE Aircrafts (
    aircraft_id SERIAL PRIMARY KEY,
    model VARCHAR(100),
    manufacturer VARCHAR(100),
    capacity INT
);

CREATE TABLE Flights (
    id SERIAL PRIMARY KEY,
    aircraft_id INT,
    departure_airport_id INT,
    arrival_airport_id INT,
    departure_time TIMESTAMP,
    arrival_time TIMESTAMP,
    distance INT,
    status VARCHAR(50),
    economy_price DECIMAL(10, 2),
    premium_economy_price DECIMAL(10, 2),
    business_price DECIMAL(10, 2),
    first_class_price DECIMAL(10, 2),
    economy_seats INT,
    premium_economy_seats INT,
    business_seats INT,
    first_class_seats INT,
    total_economy_seats INT,
    total_premium_economy_seats INT,
    total_business_seats INT,
    total_first_class_seats INT,
    FOREIGN KEY (aircraft_id) REFERENCES Aircrafts(aircraft_id),
    FOREIGN KEY (departure_airport_id) REFERENCES Airports(airport_id),
    FOREIGN KEY (arrival_airport_id) REFERENCES Airports(airport_id)
);

-- Insert Airports data
INSERT INTO Airports (name, city, IATA_code, ICAO_code) VALUES
('Indira Gandhi International Airport', 'Delhi', 'DEL', 'VIDP'),
('Chhatrapati Shivaji International Airport', 'Mumbai', 'BOM', 'VABB'),
('Kempegowda International Airport', 'Bangalore', 'BLR', 'VOBL'),
('Chennai International Airport', 'Chennai', 'MAA', 'VOMM'),
('Netaji Subhas Chandra Bose International Airport', 'Kolkata', 'CCU', 'VECC'),
('Rajiv Gandhi International Airport', 'Hyderabad', 'HYD', 'VOHS'),
('Sardar Vallabhbhai Patel International Airport', 'Ahmedabad', 'AMD', 'VAAH'),
('Cochin International Airport', 'Kochi', 'COK', 'VOCI'),
('Pune Airport', 'Pune', 'PNQ', 'VAPO'),
('Jaipur International Airport', 'Jaipur', 'JAI', 'VIJP');

-- Insert Aircrafts data
INSERT INTO Aircrafts (model, manufacturer, capacity) VALUES
('A320neo', 'Airbus', 180),
('A321neo', 'Airbus', 220),
('737-800', 'Boeing', 189),
('737 MAX 8', 'Boeing', 200),
('A319', 'Airbus', 156);

-- Insert Flights data with realistic schedules
INSERT INTO Flights (aircraft_id, departure_airport_id, arrival_airport_id, departure_time, arrival_time, distance, status, economy_price, premium_economy_price, business_price, first_class_price, economy_seats, premium_economy_seats, business_seats, first_class_seats, total_economy_seats, total_premium_economy_seats, total_business_seats, total_first_class_seats) VALUES
-- Delhi to Mumbai routes
(1, 1, 2, '2023-10-15 06:00:00', '2023-10-15 08:00:00', 1148, 'Scheduled', 100.0, 200.0, 300.0, 400.0, 100, 50, 20, 10, 100, 50, 20, 10),
(2, 1, 2, '2023-10-15 10:00:00', '2023-10-15 12:00:00', 1148, 'Scheduled', 110.0, 210.0, 310.0, 410.0, 110, 55, 22, 11, 110, 55, 22, 11),
(3, 1, 2, '2023-10-15 14:00:00', '2023-10-15 16:00:00', 1148, 'Scheduled', 120.0, 220.0, 320.0, 420.0, 120, 60, 24, 12, 120, 60, 24, 12),
(4, 1, 2, '2023-10-15 18:00:00', '2023-10-15 20:00:00', 1148, 'Scheduled', 130.0, 230.0, 330.0, 430.0, 130, 65, 26, 13, 130, 65, 26, 13),

-- Mumbai to Delhi routes
(1, 2, 1, '2023-10-15 08:30:00', '2023-10-15 10:30:00', 1148, 'Scheduled', 105.0, 205.0, 305.0, 405.0, 105, 52, 21, 10, 105, 52, 21, 10),
(2, 2, 1, '2023-10-15 12:30:00', '2023-10-15 14:30:00', 1148, 'Scheduled', 115.0, 215.0, 315.0, 415.0, 115, 57, 23, 11, 115, 57, 23, 11),
(3, 2, 1, '2023-10-15 16:30:00', '2023-10-15 18:30:00', 1148, 'Scheduled', 125.0, 225.0, 325.0, 425.0, 125, 62, 25, 12, 125, 62, 25, 12),

-- Bangalore to Delhi routes
(4, 3, 1, '2023-10-15 07:00:00', '2023-10-15 09:30:00', 1740, 'Scheduled', 150.0, 250.0, 350.0, 450.0, 150, 75, 30, 15, 150, 75, 30, 15),
(5, 3, 1, '2023-10-15 15:00:00', '2023-10-15 17:30:00', 1740, 'Scheduled', 160.0, 260.0, 360.0, 460.0, 160, 80, 32, 16, 160, 80, 32, 16),

-- Delhi to Bangalore routes
(1, 1, 3, '2023-10-15 11:00:00', '2023-10-15 13:30:00', 1740, 'Scheduled', 170.0, 270.0, 370.0, 470.0, 170, 85, 34, 17, 170, 85, 34, 17),
(2, 1, 3, '2023-10-15 19:00:00', '2023-10-15 21:30:00', 1740, 'Scheduled', 180.0, 280.0, 380.0, 480.0, 180, 90, 36, 18, 180, 90, 36, 18),

-- Mumbai to Bangalore routes
(3, 2, 3, '2023-10-15 08:00:00', '2023-10-15 09:45:00', 842, 'Scheduled', 120.0, 220.0, 320.0, 420.0, 120, 60, 24, 12, 120, 60, 24, 12),
(4, 2, 3, '2023-10-15 16:00:00', '2023-10-15 17:45:00', 842, 'Scheduled', 130.0, 230.0, 330.0, 430.0, 130, 65, 26, 13, 130, 65, 26, 13),

-- Bangalore to Mumbai routes
(5, 3, 2, '2023-10-15 10:15:00', '2023-10-15 12:00:00', 842, 'Scheduled', 140.0, 240.0, 340.0, 440.0, 140, 70, 28, 14, 140, 70, 28, 14),
(1, 3, 2, '2023-10-15 18:15:00', '2023-10-15 20:00:00', 842, 'Scheduled', 150.0, 250.0, 350.0, 450.0, 150, 75, 30, 15, 150, 75, 30, 15),

-- Chennai to Mumbai routes
(2, 4, 2, '2023-10-15 07:30:00', '2023-10-15 09:30:00', 1033, 'Scheduled', 110.0, 210.0, 310.0, 410.0, 110, 55, 22, 11, 110, 55, 22, 11),
(3, 4, 2, '2023-10-15 15:30:00', '2023-10-15 17:30:00', 1033, 'Scheduled', 120.0, 220.0, 320.0, 420.0, 120, 60, 24, 12, 120, 60, 24, 12),

-- Mumbai to Chennai routes
(4, 2, 4, '2023-10-15 10:00:00', '2023-10-15 12:00:00', 1033, 'Scheduled', 130.0, 230.0, 330.0, 430.0, 130, 65, 26, 13, 130, 65, 26, 13),
(5, 2, 4, '2023-10-15 18:00:00', '2023-10-15 20:00:00', 1033, 'Scheduled', 140.0, 240.0, 340.0, 440.0, 140, 70, 28, 14, 140, 70, 28, 14),

-- Delhi to Kolkata routes
(1, 1, 5, '2023-10-15 06:30:00', '2023-10-15 08:45:00', 1305, 'Scheduled', 150.0, 250.0, 350.0, 450.0, 150, 75, 30, 15, 150, 75, 30, 15),
(2, 1, 5, '2023-10-15 14:30:00', '2023-10-15 16:45:00', 1305, 'Scheduled', 160.0, 260.0, 360.0, 460.0, 160, 80, 32, 16, 160, 80, 32, 16),

-- Kolkata to Delhi routes
(3, 5, 1, '2023-10-15 09:15:00', '2023-10-15 11:30:00', 1305, 'Scheduled', 170.0, 270.0, 370.0, 470.0, 170, 85, 34, 17, 170, 85, 34, 17),
(4, 5, 1, '2023-10-15 17:15:00', '2023-10-15 19:30:00', 1305, 'Scheduled', 180.0, 280.0, 380.0, 480.0, 180, 90, 36, 18, 180, 90, 36, 18),

-- Hyderabad to Mumbai routes
(5, 6, 2, '2023-10-15 08:00:00', '2023-10-15 09:30:00', 631, 'Scheduled', 110.0, 210.0, 310.0, 410.0, 110, 55, 22, 11, 110, 55, 22, 11),
(1, 6, 2, '2023-10-15 16:00:00', '2023-10-15 17:30:00', 631, 'Scheduled', 120.0, 220.0, 320.0, 420.0, 120, 60, 24, 12, 120, 60, 24, 12),

-- Mumbai to Hyderabad routes
(2, 2, 6, '2023-10-15 10:00:00', '2023-10-15 11:30:00', 631, 'Scheduled', 130.0, 230.0, 330.0, 430.0, 130, 65, 26, 13, 130, 65, 26, 13),
(3, 2, 6, '2023-10-15 18:00:00', '2023-10-15 19:30:00', 631, 'Scheduled', 140.0, 240.0, 340.0, 440.0, 140, 70, 28, 14, 140, 70, 28, 14),

-- Ahmedabad to Delhi routes
(4, 7, 1, '2023-10-15 07:00:00', '2023-10-15 08:30:00', 776, 'Scheduled', 150.0, 250.0, 350.0, 450.0, 150, 75, 30, 15, 150, 75, 30, 15),
(5, 7, 1, '2023-10-15 15:00:00', '2023-10-15 16:30:00', 776, 'Scheduled', 160.0, 260.0, 360.0, 460.0, 160, 80, 32, 16, 160, 80, 32, 16),

-- Delhi to Ahmedabad routes
(1, 1, 7, '2023-10-15 09:00:00', '2023-10-15 10:30:00', 776, 'Scheduled', 170.0, 270.0, 370.0, 470.0, 170, 85, 34, 17, 170, 85, 34, 17),
(2, 1, 7, '2023-10-15 17:00:00', '2023-10-15 18:30:00', 776, 'Scheduled', 180.0, 280.0, 380.0, 480.0, 180, 90, 36, 18, 180, 90, 36, 18),

-- Additional routes between other cities
(3, 8, 3, '2023-10-15 08:00:00', '2023-10-15 09:30:00', 465, 'Scheduled', 100.0, 200.0, 300.0, 400.0, 100, 50, 20, 10, 100, 50, 20, 10),
(4, 3, 8, '2023-10-15 10:00:00', '2023-10-15 11:30:00', 465, 'Scheduled', 110.0, 210.0, 310.0, 410.0, 110, 55, 22, 11, 110, 55, 22, 11),
(5, 9, 2, '2023-10-15 07:30:00', '2023-10-15 08:30:00', 149, 'Scheduled', 120.0, 220.0, 320.0, 420.0, 120, 60, 24, 12, 120, 60, 24, 12),
(1, 2, 9, '2023-10-15 09:00:00', '2023-10-15 10:00:00', 149, 'Scheduled', 130.0, 230.0, 330.0, 430.0, 130, 65, 26, 13, 130, 65, 26, 13),
(2, 10, 1, '2023-10-15 08:00:00', '2023-10-15 09:30:00', 241, 'Scheduled', 140.0, 240.0, 340.0, 440.0, 140, 70, 28, 14, 140, 70, 28, 14),
(3, 1, 10, '2023-10-15 10:00:00', '2023-10-15 11:30:00', 241, 'Scheduled', 150.0, 250.0, 350.0, 450.0, 150, 75, 30, 15, 150, 75, 30, 15),
(4, 6, 4, '2023-10-15 09:00:00', '2023-10-15 10:30:00', 521, 'Scheduled', 160.0, 260.0, 360.0, 460.0, 160, 80, 32, 16, 160, 80, 32, 16),
(5, 4, 6, '2023-10-15 11:00:00', '2023-10-15 12:30:00', 521, 'Scheduled', 170.0, 270.0, 370.0, 470.0, 170, 85, 34, 17, 170, 85, 34, 17);