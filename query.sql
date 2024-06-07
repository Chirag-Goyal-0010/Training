-- Retrieve All Scheduled Flights for a Specific Date
SELECT * 
FROM Flights 
WHERE DATE(departure_time) = DATE '2024-05-07';

-- Update the Status of a Flight
UPDATE Flights 
SET status = 'delayed' 
WHERE flight_id = 161;

-- Get Historical Flight Data from 1st june to 5th june
SELECT * 
FROM Flights 
WHERE DATE(departure_time) BETWEEN DATE '2024-05-01' AND DATE '2024-05-05';

-- Retrieve Passenger Details for a Specific Flight
SELECT 
    f.flight_id, 
    p.passenger_id, 
    p.first_name, 
    p.last_name, 
    p.date_of_birth, 
    p.passport_number, 
    p.nationality 
FROM 
    Bookings b 
JOIN 
    Passengers p ON p.passenger_id = b.passenger_id
JOIN 
    Flights f ON f.flight_id = b.flight_id
WHERE 
    f.flight_id = 170;

-- get me details of the flight having maximum distance
SELECT * 
FROM Flights 
WHERE distance = (
    SELECT MAX(distance) 
    FROM Flights
);

-- Details of the passenger who booked flight from delhi to Goa on 1st June
SELECT 
    a1.city AS departure_city, 
    a2.city AS arrival_city, 
    f.departure_time, 
    p.passenger_id, 
    p.first_name, 
    p.last_name, 
    p.date_of_birth, 
    p.passport_number, 
    p.nationality 
FROM 
    Flights f 
JOIN 
    Airports a1 ON a1.airport_id = f.departure_airport_id
JOIN 
    Airports a2 ON a2.airport_id = f.arrival_airport_id 
JOIN 
    Bookings b ON b.flight_id = f.flight_id
JOIN 
    Passengers p ON p.passenger_id = b.passenger_id
WHERE 
    a1.city = 'New Delhi' 
    AND a2.city = 'Mumbai' 
    AND DATE(f.departure_time) = DATE '2024-05-01';

-- Find the total num of seats in all flights
SELECT SUM(A.capacity) AS total_bookings
FROM Flights F
JOIN Aircrafts A ON F.aircraft_id = A.aircraft_id;
