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
    flight_id SERIAL PRIMARY KEY,
    aircraft_id INT,
    departure_airport_id INT,
    arrival_airport_id INT,
    departure_time TIMESTAMP,
    arrival_time TIMESTAMP,
    distance INT, 
    status VARCHAR(50),
    FOREIGN KEY (aircraft_id) REFERENCES Aircrafts(aircraft_id),
    FOREIGN KEY (departure_airport_id) REFERENCES Airports(airport_id),
    FOREIGN KEY (arrival_airport_id) REFERENCES Airports(airport_id)
);

CREATE TABLE Passengers (
    passenger_id SERIAL PRIMARY KEY,
    first_name VARCHAR(50),
    last_name VARCHAR(50),
    date_of_birth DATE,
    passport_number VARCHAR(50),
    nationality VARCHAR(50)
);

CREATE TABLE Bookings (
    booking_id SERIAL PRIMARY KEY,
    flight_id INT,
    passenger_id INT,
    booking_date TIMESTAMP,
    seat_number VARCHAR(10),
    class VARCHAR(20),
    status VARCHAR(50),
    FOREIGN KEY (flight_id) REFERENCES Flights(flight_id),
    FOREIGN KEY (passenger_id) REFERENCES Passengers(passenger_id)
);

CREATE TABLE CrewMembers (
    crew_id SERIAL PRIMARY KEY,
    first_name VARCHAR(50),
    last_name VARCHAR(50),
    role VARCHAR(50)
);

CREATE TABLE Flight_Crew (
    flight_crew_id SERIAL PRIMARY KEY,
    flight_id INT,
    crew_id INT,
    FOREIGN KEY (flight_id) REFERENCES Flights(flight_id),
    FOREIGN KEY (crew_id) REFERENCES CrewMembers(crew_id)
);

SELECT table_name
FROM information_schema.tables
WHERE table_schema = 'public';

INSERT INTO Airports (name, city, IATA_code, ICAO_code) 
VALUES 
('Indira Gandhi International Airport', 'New Delhi', 'DEL', 'VIDP'),
('Chhatrapati Shivaji Maharaj International Airport', 'Mumbai', 'BOM', 'VABB'),
('Kempegowda International Airport', 'Bangalore', 'BLR', 'VOBL'),
('Chennai International Airport', 'Chennai', 'MAA', 'VOMM'),
('Netaji Subhas Chandra Bose International Airport', 'Kolkata', 'CCU', 'VECC');

INSERT INTO Aircrafts (model, manufacturer, capacity) 
VALUES 
('Boeing 737', 'Boeing', 189),
('Airbus A320', 'Airbus', 180),
('Boeing 787 Dreamliner', 'Boeing', 242),
('Airbus A321', 'Airbus', 220),
('Embraer E190', 'Embraer', 114),
('Bombardier Q400', 'Bombardier Aerospace', 78),
('Airbus A330', 'Airbus', 287),
('Boeing 777', 'Boeing', 396),
('Boeing 737 MAX', 'Boeing', 230),
('Airbus A350', 'Airbus', 440);

INSERT INTO Flights (aircraft_id, departure_airport_id, arrival_airport_id, departure_time, arrival_time, distance, status) 
VALUES 
(1, 1, 2, '2024-05-01 08:00:00', '2024-05-01 10:00:00', 715, 'On time'),
(2, 2, 1, '2024-05-02 10:30:00', '2024-05-02 12:30:00', 715, 'Delayed'),
(3, 1, 3, '2024-05-03 09:00:00', '2024-05-03 11:30:00', 1081, 'On time'),
(4, 3, 1, '2024-05-04 12:00:00', '2024-05-04 14:30:00', 1081, 'On time'),
(5, 1, 4, '2024-05-05 07:30:00', '2024-05-05 09:30:00', 1087, 'Delayed'),
(6, 4, 1, '2024-05-06 11:00:00', '2024-05-06 13:00:00', 1087, 'On time'),
(7, 1, 5, '2024-05-07 10:00:00', '2024-05-07 12:30:00', 817, 'On time'),
(8, 5, 1, '2024-05-08 13:00:00', '2024-05-08 15:30:00', 817, 'Delayed'),
(9, 1, 2, '2024-05-09 08:30:00', '2024-05-09 10:30:00', 715, 'On time'),
(10, 2, 1, '2024-05-10 11:30:00', '2024-05-10 13:30:00', 715, 'Delayed'),
(1, 1, 3, '2024-05-11 09:30:00', '2024-05-11 11:45:00', 1081, 'On time'),
(2, 3, 1, '2024-05-12 12:30:00', '2024-05-12 15:00:00', 1081, 'On time'),
(3, 1, 4, '2024-05-13 07:45:00', '2024-05-13 09:45:00', 1087, 'Delayed'),
(4, 4, 1, '2024-05-14 11:15:00', '2024-05-14 13:15:00', 1087, 'On time'),
(5, 1, 5, '2024-05-15 10:15:00', '2024-05-15 12:45:00', 817, 'On time'),
(6, 5, 1, '2024-05-16 13:15:00', '2024-05-16 15:45:00', 817, 'Delayed'),
(7, 1, 2, '2024-05-17 08:45:00', '2024-05-17 10:45:00', 715, 'On time'),
(8, 2, 1, '2024-05-18 11:45:00', '2024-05-18 13:45:00', 715, 'Delayed'),
(9, 1, 3, '2024-05-19 09:45:00', '2024-05-19 12:00:00', 1081, 'On time'),
(10, 3, 1, '2024-05-20 12:45:00', '2024-05-20 15:15:00', 1081, 'On time'),
(1, 1, 4, '2024-05-21 08:00:00', '2024-05-21 10:00:00', 1087, 'Delayed'),
(2, 4, 1, '2024-05-22 10:30:00', '2024-05-22 12:30:00', 1087, 'On time'),
(3, 1, 5, '2024-05-23 11:00:00', '2024-05-23 13:30:00', 817, 'On time'),
(4, 5, 1, '2024-05-24 14:00:00', '2024-05-24 16:30:00', 817, 'Delayed'),
(5, 1, 2, '2024-05-25 08:15:00', '2024-05-25 10:15:00', 715, 'On time'),
(6, 2, 1, '2024-05-26 11:15:00', '2024-05-26 13:15:00', 715, 'Delayed'),
(7, 1, 3, '2024-05-27 09:15:00', '2024-05-27 11:30:00', 1081, 'On time'),
(8, 3, 1, '2024-05-28 12:15:00', '2024-05-28 14:45:00', 1081, 'On time'),
(9, 1, 4, '2024-05-29 07:30:00', '2024-05-29 09:30:00', 1087, 'Delayed'),
(10, 4, 1, '2024-05-30 11:00:00', '2024-05-30 13:00:00', 1087, 'On time'),
(1, 1, 5, '2024-05-31 10:30:00', '2024-05-31 13:00:00', 817, 'On time');


INSERT INTO Passengers (first_name, last_name, date_of_birth, passport_number, nationality)
SELECT
    substr(md5(random()::text), 1, 10) AS first_name,
    substr(md5(random()::text), 1, 10) AS last_name,
    (DATE '1950-01-01' + (random() * (DATE '2005-12-31' - DATE '1950-01-01')::int)::int)::date AS date_of_birth,
    substr(md5(random()::text), 1, 8) AS passport_number,
    (array['India', 'Pakistan', 'Bangladesh', 'Bhutan', 'Nepal', 'USA', 'China', 'Maldives'])[floor(random() * 8 + 1)::int] AS nationality
FROM
    generate_series(1, 1000);

-- Insert bookings assuming 10% capacity
WITH AircraftCapacity AS (
    SELECT aircraft_id, capacity, 
           CEIL(capacity * 0.1) AS booked_seats
    FROM Aircrafts
)
INSERT INTO Bookings (flight_id, passenger_id, booking_date, seat_number, class, status)
SELECT
    f.flight_id,
    p.passenger_id,
    NOW(),  -- Assuming booking date is the current timestamp
    'A' || FLOOR(random() * 6 + 1)::int || FLOOR(random() * 30 + 1),  -- Random seat number
    (array['Economy', 'Business', 'First'])[FLOOR(random() * 3 + 1)::int],  -- Random class
    'Confirmed'
FROM Flights f
JOIN AircraftCapacity ac ON f.aircraft_id = ac.aircraft_id,
LATERAL (
    SELECT passenger_id
    FROM Passengers
    ORDER BY random()
    LIMIT ac.booked_seats
) p;

UPDATE Bookings AS b
SET booking_date = f.departure_time - INTERVAL '1 day' * FLOOR(random() * 15)::int
FROM Flights AS f
WHERE b.flight_id = f.flight_id;




INSERT INTO CrewMembers (first_name, last_name, role)
SELECT
    substr(md5(random()::text), 1, 10) AS first_name,
    substr(md5(random()::text), 1, 10) AS last_name,
    (array['Pilot', 'Co-Pilot', 'Flight Attendant', 'Engineer', 'Navigator'])[floor(random() * 5 + 1)::int] AS role
FROM
    generate_series(1, 50);

INSERT INTO Flight_Crew (flight_id, crew_id)
SELECT
    f.flight_id,
    c.crew_id
FROM
    Flights f
CROSS JOIN LATERAL (
    SELECT crew_id
    FROM CrewMembers
    ORDER BY random()
    LIMIT 1
) c
LIMIT 50;

-- INSERT INTO Bookings (flight_id, passenger_id, booking_date, seat_number, class, status)
-- SELECT 
--     floor(random() * 31 + 151) AS flight_id,
--     generate_series(1, 1000) AS passenger_id,
--     '2024-06-01'::TIMESTAMP + random() * ('2024-06-30'::TIMESTAMP - '2024-06-01'::TIMESTAMP) AS booking_date,
--     'Seat ' || floor(random() * 200 + 1)::INT AS seat_number,
--     CASE 
--         WHEN random() < 0.7 THEN 'Economy'
--         ELSE 'Business'
--     END AS class,
--     CASE 
--         WHEN random() < 0.9 THEN 'Confirmed'
--         ELSE 'Pending'
--     END AS status
-- FROM 
--     generate_series(1, 1000)
-- WHERE
--     floor(random() * 31 + 151) <= 181;