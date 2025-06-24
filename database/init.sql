ALTER SCHEMA public OWNER TO postgres;

SET default_tablespace = '';
SET default_table_access_method = heap;

CREATE TABLE public.aircrafts (
    aircraft_id integer NOT NULL,
    model character varying(100),
    manufacturer character varying(100),
    capacity integer
);

ALTER TABLE public.aircrafts OWNER TO postgres;

CREATE SEQUENCE public.aircrafts_aircraft_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.aircrafts_aircraft_id_seq OWNER TO postgres;
ALTER SEQUENCE public.aircrafts_aircraft_id_seq OWNED BY public.aircrafts.aircraft_id;

CREATE TABLE public.airports (
    airport_id integer NOT NULL,
    name character varying(100),
    city character varying(50),
    iata_code character varying(10),
    icao_code character varying(10)
);

ALTER TABLE public.airports OWNER TO postgres;

CREATE SEQUENCE public.airports_airport_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.airports_airport_id_seq OWNER TO postgres;
ALTER SEQUENCE public.airports_airport_id_seq OWNED BY public.airports.airport_id;

CREATE TABLE public.bookings (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    user_id bigint,
    flight_id bigint,
    booking_date timestamp with time zone,
    status text,
    seats bigint,
    is_premium boolean,
    total_price numeric,
    travel_class text
);

ALTER TABLE public.bookings OWNER TO postgres;

CREATE SEQUENCE public.bookings_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.bookings_id_seq OWNER TO postgres;
ALTER SEQUENCE public.bookings_id_seq OWNED BY public.bookings.id;

CREATE TABLE public.flights (
    id integer NOT NULL,
    aircraft_id integer,
    departure_airport_id integer,
    arrival_airport_id integer,
    departure_time timestamp without time zone,
    arrival_time timestamp without time zone,
    distance integer,
    status character varying(50),
    economy_price numeric(10,2),
    premium_economy_price numeric(10,2),
    business_price numeric(10,2),
    first_class_price numeric(10,2),
    economy_seats bigint,
    premium_economy_seats bigint,
    business_seats bigint,
    first_class_seats bigint,
    total_economy_seats bigint,
    total_premium_economy_seats bigint,
    total_business_seats bigint,
    total_first_class_seats bigint,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    origin text,
    destination text
);

ALTER TABLE public.flights OWNER TO postgres;

CREATE SEQUENCE public.flights_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.flights_id_seq OWNER TO postgres;
ALTER SEQUENCE public.flights_id_seq OWNED BY public.flights.id;

CREATE TABLE public.travellers (
    id bigint NOT NULL,
    booking_id bigint,
    title text,
    first_name text,
    last_name text,
    dob timestamp with time zone,
    nationality text
);

ALTER TABLE public.travellers OWNER TO postgres;

CREATE SEQUENCE public.travellers_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.travellers_id_seq OWNER TO postgres;
ALTER SEQUENCE public.travellers_id_seq OWNED BY public.travellers.id;

CREATE TABLE public.users (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    email text NOT NULL,
    password text,
    name text,
    is_admin boolean
);

ALTER TABLE public.users OWNER TO postgres;

CREATE SEQUENCE public.users_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.users_id_seq OWNER TO postgres;
ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;

ALTER TABLE ONLY public.aircrafts ALTER COLUMN aircraft_id SET DEFAULT nextval('public.aircrafts_aircraft_id_seq'::regclass);
ALTER TABLE ONLY public.airports ALTER COLUMN airport_id SET DEFAULT nextval('public.airports_airport_id_seq'::regclass);
ALTER TABLE ONLY public.bookings ALTER COLUMN id SET DEFAULT nextval('public.bookings_id_seq'::regclass);
ALTER TABLE ONLY public.flights ALTER COLUMN id SET DEFAULT nextval('public.flights_id_seq'::regclass);
ALTER TABLE ONLY public.travellers ALTER COLUMN id SET DEFAULT nextval('public.travellers_id_seq'::regclass);
ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);

INSERT INTO public.aircrafts (aircraft_id, model, manufacturer, capacity) VALUES
(1, 'Boeing 737', 'Boeing', 180),
(2, 'Airbus A320', 'Airbus', 160),
(3, 'Boeing 777', 'Boeing', 396),
(4, 'Airbus A380', 'Airbus', 555),
(5, 'Embraer 190', 'Embraer', 100),
(6, 'Bombardier Q400', 'Bombardier', 78),
(7, 'Boeing 787 Dreamliner', 'Boeing', 242),
(8, 'Airbus A350', 'Airbus', 325),
(9, 'ATR 72', 'ATR', 70),
(10, 'Cessna 208 Caravan', 'Cessna', 12);


INSERT INTO public.airports (airport_id, name, city, iata_code, icao_code) VALUES
(1, 'Indira Gandhi International Airport', 'Delhi', 'DEL', 'VIDP'),
(2, 'Chhatrapati Shivaji International Airport', 'Mumbai', 'BOM', 'VABB'),
(3, 'Kempegowda International Airport', 'Bangalore', 'BLR', 'VOBL'),
(4, 'Chennai International Airport', 'Chennai', 'MAA', 'VOMM'),
(5, 'Netaji Subhas Chandra Bose International Airport', 'Kolkata', 'CCU', 'VECC'),
(6, 'Rajiv Gandhi International Airport', 'Hyderabad', 'HYD', 'VOHS'),
(7, 'Sardar Vallabhbhai Patel International Airport', 'Ahmedabad', 'AMD', 'VAAH'),
(8, 'Cochin International Airport', 'Kochi', 'COK', 'VOCI'),
(9, 'Pune Airport', 'Pune', 'PNQ', 'VAPO'),
(10, 'Jaipur International Airport', 'Jaipur', 'JAI', 'VIJP');

INSERT INTO public.bookings (
    id, created_at, updated_at, deleted_at, user_id, flight_id, booking_date, status, seats, is_premium, total_price, travel_class
) VALUES
(1, '2025-06-17 16:51:01.192154+05:30', '2025-06-17 16:51:01.192154+05:30', NULL, 1, 40, '2025-06-17 16:51:01.191638+05:30', 'Confirmed', 50, false, 10000, 'PremiumEconomy'),
(2, '2025-06-17 16:51:42.341966+05:30', '2025-06-17 16:51:42.341966+05:30', NULL, 1, 41, '2025-06-17 16:51:42.341966+05:30', 'Confirmed', 10, false, 1000, 'Economy'),
(3, '2025-06-17 16:51:50.626017+05:30', '2025-06-17 16:51:50.626017+05:30', NULL, 1, 42, '2025-06-17 16:51:50.626017+05:30', 'Confirmed', 15, false, 3000, 'PremiumEconomy'),
(4, '2025-06-17 16:51:58.742396+05:30', '2025-06-17 16:51:58.742396+05:30', NULL, 1, 43, '2025-06-17 16:51:58.742396+05:30', 'Confirmed', 6, false, 2400, 'FirstClass'),
(5, '2025-06-17 16:52:07.797883+05:30', '2025-06-17 16:52:07.797883+05:30', NULL, 1, 44, '2025-06-17 16:52:07.797883+05:30', 'Confirmed', 10, false, 2990, 'Business'),
(6, '2025-06-17 17:04:04.99755+05:30', '2025-06-17 17:04:04.99755+05:30', NULL, 2, 40, '2025-06-17 17:04:04.997031+05:30', 'Confirmed', 10, false, 1000, 'Economy'),
(7, '2025-06-17 17:04:17.243831+05:30', '2025-06-17 17:04:17.243831+05:30', NULL, 2, 41, '2025-06-17 17:04:17.243326+05:30', 'Confirmed', 4, false, 800, 'PremiumEconomy'),
(8, '2025-06-17 17:04:25.392322+05:30', '2025-06-17 17:04:25.392322+05:30', NULL, 2, 43, '2025-06-17 17:04:25.392322+05:30', 'Confirmed', 100, false, 9500, 'Economy'),
(9, '2025-06-17 17:04:34.807091+05:30', '2025-06-17 17:04:34.807091+05:30', NULL, 2, 44, '2025-06-17 17:04:34.807091+05:30', 'Confirmed', 3, false, 897, 'Business'),
(10, '2025-06-18 11:29:14.214551+05:30', '2025-06-18 11:29:14.214551+05:30', NULL, 2, 40, '2025-06-18 11:29:14.214551+05:30', 'Confirmed', 10, false, 1000, 'Economy'),
(11, '2025-06-18 11:29:35.851431+05:30', '2025-06-18 11:29:35.851431+05:30', NULL, 2, 40, '2025-06-18 11:29:35.851431+05:30', 'Confirmed', 10, false, 2990, 'Business'),
(12, '2025-06-18 11:30:07.67769+05:30', '2025-06-18 11:30:07.67769+05:30', NULL, 2, 42, '2025-06-18 11:30:07.67769+05:30', 'Confirmed', 5, false, 1495, 'Business'),
(13, '2025-06-18 11:33:43.18302+05:30', '2025-06-18 11:33:43.18302+05:30', NULL, 2, 44, '2025-06-18 11:33:43.182505+05:30', 'Confirmed', 6, false, 1794, 'Business'),
(14, '2025-06-18 11:56:35.159726+05:30', '2025-06-18 11:56:35.159726+05:30', NULL, 2, 44, '2025-06-18 11:56:35.15901+05:30', 'Confirmed', 2, false, 190, 'Economy'),
(15, '2025-06-18 12:12:49.908653+05:30', '2025-06-18 12:12:49.908653+05:30', NULL, 2, 53, '2025-06-18 12:12:49.907633+05:30', 'Confirmed', 3, false, 1200, 'FirstClass');

INSERT INTO public.flights (id, aircraft_id, departure_airport_id, arrival_airport_id, departure_time, arrival_time, distance, status, economy_price, premium_economy_price, business_price, first_class_price, economy_seats, premium_economy_seats, business_seats, first_class_seats, total_economy_seats, total_premium_economy_seats, total_business_seats, total_first_class_seats, created_at, updated_at, deleted_at, origin, destination) VALUES
(1, 3, 9, 1, '2024-06-28 10:00:00', '2024-06-28 12:00:00', 1552, 'Scheduled', 332.77, 529.74, 679.58, 853.05, 116, 66, 13, 6, 116, 66, 13, 6, '2025-06-19 10:36:03.412739+05:30', '2025-06-19 10:36:03.412739+05:30', NULL, 'Pune', 'Delhi'),
(2, 2, 10, 4, '2024-06-26 08:00:00', '2024-06-26 11:00:00', 1009, 'Scheduled', 166.31, 340.52, 455.91, 583.11, 101, 76, 31, 19, 101, 76, 31, 19, '2025-06-19 10:36:03.412739+05:30', '2025-06-19 10:36:03.412739+05:30', NULL, 'Jaipur', 'Chennai'),
(3, 1, 2, 9, '2024-06-26 10:00:00', '2024-06-26 14:30:00', 193, 'Scheduled', 483.13, 588.01, 727.65, 841.31, 162, 43, 26, 20, 162, 43, 26, 20, '2025-06-19 10:36:03.412739+05:30', '2025-06-19 10:36:03.412739+05:30', NULL, 'Mumbai', 'Pune'),
(4, 5, 7, 2, '2024-06-28 06:00:00', '2024-06-28 10:30:00', 1414, 'Scheduled', 224.41, 418.05, 586.82, 696.57, 139, 60, 36, 10, 139, 60, 36, 10, '2025-06-19 10:36:03.412739+05:30', '2025-06-19 10:36:03.412739+05:30', NULL, 'Ahmedabad', 'Mumbai'),
(5, 1, 4, 7, '2024-06-27 09:00:00', '2024-06-27 13:00:00', 363, 'Scheduled', 499.47, 639.78, 816.72, 932.96, 134, 82, 19, 19, 134, 82, 19, 19, '2025-06-19 10:36:03.412739+05:30', '2025-06-19 10:36:03.412739+05:30', NULL, 'Chennai', 'Ahmedabad'),
(6, 1, 4, 10, '2024-06-26 09:00:00', '2024-06-26 11:00:00', 562, 'Scheduled', 287.04, 481.24, 602.96, 714.77, 126, 87, 20, 11, 126, 87, 20, 11, '2025-06-19 10:36:03.412739+05:30', '2025-06-19 10:36:03.412739+05:30', NULL, 'Chennai', 'Jaipur'),
(7, 2, 4, 10, '2024-06-28 09:00:00', '2024-06-28 11:00:00', 197, 'Scheduled', 251.64, 376.51, 570.35, 743.77, 115, 46, 35, 19, 115, 46, 35, 19, '2025-06-19 10:36:03.412739+05:30', '2025-06-19 10:36:03.412739+05:30', NULL, 'Chennai', 'Jaipur'),
(8, 5, 5, 7, '2024-06-28 07:00:00', '2024-06-28 09:00:00', 1440, 'Scheduled', 321.23, 509.07, 664.31, 845.53, 96, 70, 22, 13, 96, 70, 22, 13, '2025-06-19 10:36:03.412739+05:30', '2025-06-19 10:36:03.412739+05:30', NULL, 'Kolkata', 'Ahmedabad'),
(9, 3, 9, 7, '2024-06-26 07:00:00', '2024-06-26 11:30:00', 693, 'Scheduled', 340.76, 526.4, 725.97, 901.31, 99, 30, 10, 18, 99, 30, 10, 18, '2025-06-19 10:36:03.412739+05:30', '2025-06-19 10:36:03.412739+05:30', NULL, 'Pune', 'Ahmedabad'),
(10, 4, 2, 7, '2024-06-28 10:00:00', '2024-06-28 12:30:00', 1396, 'Scheduled', 204.88, 343.95, 514.85, 714.08, 93, 33, 21, 6, 93, 33, 21, 6, '2025-06-19 10:36:03.412739+05:30', '2025-06-19 10:36:03.412739+05:30', NULL, 'Mumbai', 'Ahmedabad'),
(11, 2, 8, 1, '2024-06-26 10:00:00', '2024-06-26 12:00:00', 321, 'Scheduled', 208.92, 400.9, 507.42, 678.43, 95, 84, 14, 19, 95, 84, 14, 19, '2025-06-19 10:36:03.412739+05:30', '2025-06-19 10:36:03.412739+05:30', NULL, 'Kochi', 'Delhi'),
(12, 1, 7, 1, '2024-06-27 08:00:00', '2024-06-27 10:30:00', 1216, 'Scheduled', 449.03, 646.63, 747.63, 914.88, 154, 88, 30, 20, 154, 88, 30, 20, '2025-06-19 10:36:03.412739+05:30', '2025-06-19 10:36:03.412739+05:30', NULL, 'Ahmedabad', 'Delhi'),
(13, 5, 10, 9, '2024-06-27 10:00:00', '2024-06-27 14:00:00', 1057, 'Scheduled', 450.71, 586.7, 775.92, 880.05, 101, 87, 17, 17, 101, 87, 17, 17, '2025-06-19 10:36:03.412739+05:30', '2025-06-19 10:36:03.412739+05:30', NULL, 'Jaipur', 'Pune'),
(14, 2, 8, 3, '2024-06-26 07:00:00', '2024-06-26 09:00:00', 1541, 'Scheduled', 335.86, 528.09, 646.39, 830.2, 119, 51, 30, 9, 119, 51, 30, 9, '2025-06-19 10:36:03.412739+05:30', '2025-06-19 10:36:03.412739+05:30', NULL, 'Kochi', 'Bangalore'),
(15, 3, 10, 8, '2024-06-27 07:00:00', '2024-06-27 10:30:00', 203, 'Scheduled', 251.96, 447.0, 626.34, 794.28, 180, 84, 22, 9, 180, 84, 22, 9, '2025-06-19 10:36:03.412739+05:30', '2025-06-19 10:36:03.412739+05:30', NULL, 'Jaipur', 'Kochi'),
(16, 2, 8, 3, '2024-06-27 07:00:00', '2024-06-27 09:30:00', 1807, 'Scheduled', 187.3, 372.68, 525.7, 705.98, 80, 74, 40, 8, 80, 74, 40, 8, '2025-06-19 10:36:03.412739+05:30', '2025-06-19 10:36:03.412739+05:30', NULL, 'Kochi', 'Bangalore'),
(17, 2, 3, 2, '2024-06-29 06:00:00', '2024-06-29 08:00:00', 1874, 'Scheduled', 495.5, 664.52, 837.68, 983.34, 87, 56, 29, 6, 87, 56, 29, 6, '2025-06-19 10:36:03.412739+05:30', '2025-06-19 10:36:03.412739+05:30', NULL, 'Bangalore', 'Mumbai'),
(18, 3, 3, 1, '2024-06-26 06:00:00', '2024-06-26 09:30:00', 859, 'Scheduled', 201.23, 380.24, 571.89, 686.58, 142, 30, 33, 5, 142, 30, 33, 5, '2025-06-19 10:36:03.412739+05:30', '2025-06-19 10:36:03.412739+05:30', NULL, 'Bangalore', 'Delhi'),
(19, 1, 1, 9, '2024-06-28 07:00:00', '2024-06-28 11:30:00', 266, 'Scheduled', 208.04, 405.09, 586.25, 696.99, 129, 42, 29, 11, 129, 42, 29, 11, '2025-06-19 10:36:03.412739+05:30', '2025-06-19 10:36:03.412739+05:30', NULL, 'Delhi', 'Pune'),
(20, 1, 10, 6, '2024-06-26 08:00:00', '2024-06-26 10:00:00', 735, 'Scheduled', 181.9, 346.05, 521.69, 719.83, 134, 83, 16, 17, 134, 83, 16, 17, '2025-06-19 10:36:03.412739+05:30', '2025-06-19 10:36:03.412739+05:30', NULL, 'Jaipur', 'Hyderabad'),
(21, 1, 8, 2, '2024-06-29 09:00:00', '2024-06-29 12:00:00', 1478, 'Scheduled', 229.95, 367.25, 555.95, 690.78, 155, 35, 23, 9, 155, 35, 23, 9, '2025-06-19 10:36:03.412739+05:30', '2025-06-19 10:36:03.412739+05:30', NULL, 'Kochi', 'Mumbai'),
(22, 1, 8, 5, '2024-06-27 08:00:00', '2024-06-27 10:00:00', 791, 'Scheduled', 468.32, 574.7, 740.7, 883.58, 167, 30, 27, 19, 167, 30, 27, 19, '2025-06-19 10:36:03.412739+05:30', '2025-06-19 10:36:03.412739+05:30', NULL, 'Kochi', 'Kolkata'),
(23, 4, 10, 3, '2024-06-28 09:00:00', '2024-06-28 13:30:00', 1097, 'Scheduled', 263.66, 388.4, 579.71, 729.35, 154, 35, 33, 7, 154, 35, 33, 7, '2025-06-19 10:36:03.412739+05:30', '2025-06-19 10:36:03.412739+05:30', NULL, 'Jaipur', 'Bangalore'),
(24, 2, 2, 7, '2024-06-28 09:00:00', '2024-06-28 13:00:00', 1241, 'Scheduled', 356.14, 532.95, 721.55, 857.62, 158, 68, 32, 13, 158, 68, 32, 13, '2025-06-19 10:36:03.412739+05:30', '2025-06-19 10:36:03.412739+05:30', NULL, 'Mumbai', 'Ahmedabad'),
(25, 5, 8, 3, '2024-06-26 10:00:00', '2024-06-26 14:30:00', 1862, 'Scheduled', 484.76, 676.25, 865.35, 1009.29, 138, 89, 27, 6, 138, 89, 27, 6, '2025-06-19 10:36:03.412739+05:30', '2025-06-19 10:36:03.412739+05:30', NULL, 'Kochi', 'Bangalore'),
(26, 2, 10, 5, '2024-06-28 09:00:00', '2024-06-28 13:30:00', 188, 'Scheduled', 188.27, 338.12, 453.75, 574.81, 150, 40, 12, 5, 150, 40, 12, 5, '2025-06-19 10:36:03.412739+05:30', '2025-06-19 10:36:03.412739+05:30', NULL, 'Jaipur', 'Kolkata'),
(27, 5, 2, 3, '2024-06-27 08:00:00', '2024-06-27 10:30:00', 180, 'Scheduled', 165.97, 337.9, 448.78, 568.39, 132, 47, 10, 9, 132, 47, 10, 9, '2025-06-19 10:36:03.412739+05:30', '2025-06-19 10:36:03.412739+05:30', NULL, 'Mumbai', 'Bangalore'),
(28, 2, 4, 3, '2024-06-26 06:00:00', '2024-06-26 09:00:00', 1409, 'Scheduled', 396.66, 579.63, 724.11, 837.54, 80, 31, 20, 16, 80, 31, 20, 16, '2025-06-19 10:36:03.412739+05:30', '2025-06-19 10:36:03.412739+05:30', NULL, 'Chennai', 'Bangalore'),
(29, 5, 8, 5, '2024-06-29 08:00:00', '2024-06-29 10:30:00', 880, 'Scheduled', 406.97, 598.21, 723.85, 868.57, 126, 32, 26, 20, 126, 32, 26, 20, '2025-06-19 10:36:03.412739+05:30', '2025-06-19 10:36:03.412739+05:30', NULL, 'Kochi', 'Kolkata'),
(30, 5, 5, 6, '2024-06-29 07:00:00', '2024-06-29 09:30:00', 1418, 'Scheduled', 258.64, 389.6, 547.19, 651.96, 107, 71, 19, 9, 107, 71, 19, 9, '2025-06-19 10:36:03.412739+05:30', '2025-06-19 10:36:03.412739+05:30', NULL, 'Kolkata', 'Hyderabad'),
(31, 2, 4, 5, '2024-06-26 09:00:00', '2024-06-26 11:00:00', 316, 'Scheduled', 206.91, 340.34, 470.5, 660.2, 118, 53, 33, 14, 118, 53, 33, 14, '2025-06-19 10:36:03.412739+05:30', '2025-06-19 10:36:03.412739+05:30', NULL, 'Chennai', 'Kolkata'),
(32, 5, 5, 2, '2024-06-27 08:00:00', '2024-06-27 11:00:00', 904, 'Scheduled', 394.3, 549.08, 681.54, 859.54, 174, 59, 25, 5, 174, 59, 25, 5, '2025-06-19 10:36:03.412739+05:30', '2025-06-19 10:36:03.412739+05:30', NULL, 'Kolkata', 'Mumbai'),
(33, 3, 1, 7, '2024-06-28 09:00:00', '2024-06-28 12:00:00', 1067, 'Scheduled', 340.63, 467.69, 627.64, 801.23, 170, 85, 39, 17, 170, 85, 39, 17, '2025-06-19 10:36:03.412739+05:30', '2025-06-19 10:36:03.412739+05:30', NULL, 'Delhi', 'Ahmedabad'),
(34, 5, 7, 2, '2024-06-28 09:00:00', '2024-06-28 12:30:00', 1432, 'Scheduled', 296.33, 398.88, 582.33, 685.78, 140, 38, 26, 18, 140, 38, 26, 18, '2025-06-19 10:36:03.412739+05:30', '2025-06-19 10:36:03.412739+05:30', NULL, 'Ahmedabad', 'Mumbai'),
(35, 5, 9, 10, '2024-06-26 07:00:00', '2024-06-26 09:30:00', 1331, 'Scheduled', 462.07, 576.74, 764.02, 935.53, 158, 40, 28, 11, 158, 40, 28, 11, '2025-06-19 10:36:03.412739+05:30', '2025-06-19 10:36:03.412739+05:30', NULL, 'Pune', 'Jaipur'),
(36, 4, 6, 3, '2024-06-27 06:00:00', '2024-06-27 08:00:00', 653, 'Scheduled', 164.56, 289.46, 394.92, 543.04, 162, 52, 20, 13, 162, 52, 20, 13, '2025-06-19 10:36:03.412739+05:30', '2025-06-19 10:36:03.412739+05:30', NULL, 'Hyderabad', 'Bangalore'),
(37, 1, 8, 5, '2024-06-27 07:00:00', '2024-06-27 09:00:00', 1130, 'Scheduled', 242.57, 382.64, 552.16, 662.95, 133, 90, 28, 14, 133, 90, 28, 14, '2025-06-19 10:36:03.412739+05:30', '2025-06-19 10:36:03.412739+05:30', NULL, 'Kochi', 'Kolkata'),
(38, 5, 3, 2, '2024-06-27 08:00:00', '2024-06-27 11:00:00', 687, 'Scheduled', 249.73, 355.72, 498.26, 693.05, 81, 72, 24, 10, 81, 72, 24, 10, '2025-06-19 10:36:03.412739+05:30', '2025-06-19 10:36:03.412739+05:30', NULL, 'Bangalore', 'Mumbai'),
(39, 5, 8, 1, '2024-06-26 08:00:00', '2024-06-26 11:00:00', 1737, 'Scheduled', 284.2, 467.01, 659.67, 826.19, 138, 78, 32, 5, 138, 78, 32, 5, '2025-06-19 10:36:03.412739+05:30', '2025-06-19 10:36:03.412739+05:30', NULL, 'Kochi', 'Delhi'),
(40, 1, 3, 1, '2024-06-26 06:00:00', '2024-06-26 10:00:00', 1477, 'Scheduled', 246.87, 398.85, 579.71, 729.52, 174, 62, 30, 7, 174, 62, 30, 7, '2025-06-19 10:36:03.412739+05:30', '2025-06-19 10:36:03.412739+05:30', NULL, 'Bangalore', 'Delhi'),
(41, 3, 9, 10, '2024-06-28 07:00:00', '2024-06-28 11:30:00', 1014, 'Scheduled', 443.72, 554.49, 682.91, 788.52, 162, 70, 12, 8, 162, 70, 12, 8, '2025-06-19 10:36:03.412739+05:30', '2025-06-19 10:36:03.412739+05:30', NULL, 'Pune', 'Jaipur'),
(42, 1, 2, 4, '2024-06-29 07:00:00', '2024-06-29 11:30:00', 242, 'Scheduled', 275.67, 464.04, 565.73, 674.45, 156, 60, 19, 15, 156, 60, 19, 15, '2025-06-19 10:36:03.412739+05:30', '2025-06-19 10:36:03.412739+05:30', NULL, 'Mumbai', 'Chennai'),
(43, 2, 6, 5, '2024-06-26 08:00:00', '2024-06-26 12:30:00', 892, 'Scheduled', 445.67, 565.87, 719.58, 918.38, 91, 76, 26, 10, 91, 76, 26, 10, '2025-06-19 10:36:03.412739+05:30', '2025-06-19 10:36:03.412739+05:30', NULL, 'Hyderabad', 'Kolkata'),
(44, 1, 4, 9, '2024-06-26 06:00:00', '2024-06-26 08:30:00', 345, 'Scheduled', 465.19, 639.7, 742.05, 896.62, 96, 41, 23, 16, 96, 41, 23, 16, '2025-06-19 10:36:03.412739+05:30', '2025-06-19 10:36:03.412739+05:30', NULL, 'Chennai', 'Pune'),
(45, 1, 6, 9, '2024-06-28 06:00:00', '2024-06-28 09:00:00', 1759, 'Scheduled', 302.72, 414.7, 529.45, 667.68, 111, 48, 30, 20, 111, 48, 30, 20, '2025-06-19 10:36:03.412739+05:30', '2025-06-19 10:36:03.412739+05:30', NULL, 'Hyderabad', 'Pune'),
(46, 5, 3, 9, '2024-06-27 06:00:00', '2024-06-27 08:00:00', 1837, 'Scheduled', 360.01, 500.26, 647.03, 767.42, 173, 78, 35, 9, 173, 78, 35, 9, '2025-06-19 10:36:03.412739+05:30', '2025-06-19 10:36:03.412739+05:30', NULL, 'Bangalore', 'Pune'),
(47, 2, 3, 8, '2024-06-28 06:00:00', '2024-06-28 09:30:00', 618, 'Scheduled', 281.59, 397.6, 594.87, 727.93, 158, 55, 23, 20, 158, 55, 23, 20, '2025-06-19 10:36:03.412739+05:30', '2025-06-19 10:36:03.412739+05:30', NULL, 'Bangalore', 'Kochi'),
(48, 2, 9, 6, '2024-06-28 10:00:00', '2024-06-28 12:00:00', 1594, 'Scheduled', 374.96, 548.6, 681.68, 789.01, 96, 69, 23, 9, 96, 69, 23, 9, '2025-06-19 10:36:03.412739+05:30', '2025-06-19 10:36:03.412739+05:30', NULL, 'Pune', 'Hyderabad'),
(49, 5, 9, 7, '2024-06-28 06:00:00', '2024-06-28 09:00:00', 1670, 'Scheduled', 422.59, 588.76, 749.91, 909.8, 136, 34, 15, 8, 136, 34, 15, 8, '2025-06-19 10:36:03.412739+05:30', '2025-06-19 10:36:03.412739+05:30', NULL, 'Pune', 'Ahmedabad'),
(50, 3, 1, 5, '2024-06-27 06:00:00', '2024-06-27 08:30:00', 737, 'Scheduled', 263.45, 379.81, 489.65, 611.26, 160, 72, 26, 9, 160, 72, 26, 9, '2025-06-19 10:36:03.412739+05:30', '2025-06-19 10:36:03.412739+05:30', NULL, 'Delhi', 'Kolkata'),
(51, 1, 1, 10, '2024-06-26 06:00:00', '2024-06-26 09:00:00', 433, 'Scheduled', 419.21, 552.19, 746.32, 937.15, 134, 45, 30, 20, 134, 45, 30, 20, '2025-06-19 10:36:03.412739+05:30', '2025-06-19 10:36:03.412739+05:30', NULL, 'Delhi', 'Jaipur'),
(52, 5, 8, 1, '2024-06-27 09:00:00', '2024-06-27 11:00:00', 416, 'Scheduled', 174.75, 280.81, 432.31, 556.39, 127, 42, 26, 20, 127, 42, 26, 20, '2025-06-19 10:36:03.412739+05:30', '2025-06-19 10:36:03.412739+05:30', NULL, 'Kochi', 'Delhi'),
(53, 2, 7, 8, '2024-06-29 08:00:00', '2024-06-29 11:30:00', 1484, 'Scheduled', 424.07, 568.6, 697.51, 876.18, 131, 78, 33, 11, 131, 78, 33, 11, '2025-06-19 10:36:03.412739+05:30', '2025-06-19 10:36:03.412739+05:30', NULL, 'Ahmedabad', 'Kochi'),
(54, 4, 6, 4, '2024-06-26 09:00:00', '2024-06-26 11:00:00', 659, 'Scheduled', 331.69, 521.33, 624.76, 767.1, 131, 51, 12, 9, 131, 51, 12, 9, '2025-06-19 10:36:03.412739+05:30', '2025-06-19 10:36:03.412739+05:30', NULL, 'Hyderabad', 'Chennai'),
(55, 1, 6, 4, '2024-06-29 10:00:00', '2024-06-29 14:00:00', 543, 'Scheduled', 277.07, 405.11, 533.24, 666.37, 134, 66, 22, 19, 134, 66, 22, 19, '2025-06-19 10:36:03.412739+05:30', '2025-06-19 10:36:03.412739+05:30', NULL, 'Hyderabad', 'Chennai'),
(56, 1, 1, 9, '2024-06-28 10:00:00', '2024-06-28 14:30:00', 907, 'Scheduled', 260.33, 373.94, 523.67, 661.19, 113, 77, 29, 19, 113, 77, 29, 19, '2025-06-19 10:36:03.412739+05:30', '2025-06-19 10:36:03.412739+05:30', NULL, 'Delhi', 'Pune'),
(57, 4, 9, 5, '2024-06-26 08:00:00', '2024-06-26 10:30:00', 1950, 'Scheduled', 430.64, 578.38, 686.69, 812.17, 108, 89, 40, 15, 108, 89, 40, 15, '2025-06-19 10:36:03.412739+05:30', '2025-06-19 10:36:03.412739+05:30', NULL, 'Pune', 'Kolkata'),
(58, 1, 4, 10, '2024-06-26 08:00:00', '2024-06-26 11:30:00', 489, 'Scheduled', 400.72, 540.9, 735.32, 918.02, 98, 52, 17, 8, 98, 52, 17, 8, '2025-06-19 10:36:03.412739+05:30', '2025-06-19 10:36:03.412739+05:30', NULL, 'Chennai', 'Jaipur'),
(59, 3, 10, 7, '2024-06-29 09:00:00', '2024-06-29 11:00:00', 1994, 'Scheduled', 251.75, 390.39, 578.52, 696.08, 122, 50, 32, 18, 122, 50, 32, 18, '2025-06-19 10:36:03.412739+05:30', '2025-06-19 10:36:03.412739+05:30', NULL, 'Jaipur', 'Ahmedabad'),
(60, 3, 9, 7, '2024-06-27 09:00:00', '2024-06-27 11:30:00', 843, 'Scheduled', 329.03, 527.32, 676.61, 868.24, 176, 54, 37, 20, 176, 54, 37, 20, '2025-06-19 10:36:03.412739+05:30', '2025-06-19 10:36:03.412739+05:30', NULL, 'Pune', 'Ahmedabad'),
(61, 3, 2, 9, '2024-06-28 10:00:00', '2024-06-28 13:30:00', 1561, 'Scheduled', 473.48, 641.93, 830.88, 995.98, 163, 71, 31, 5, 163, 71, 31, 5, '2025-06-19 10:36:03.412739+05:30', '2025-06-19 10:36:03.412739+05:30', NULL, 'Mumbai', 'Pune'),
(62, 4, 7, 9, '2024-06-28 08:00:00', '2024-06-28 11:30:00', 1262, 'Scheduled', 194.2, 317.84, 432.0, 535.51, 112, 88, 35, 18, 112, 88, 35, 18, '2025-06-19 10:36:03.412739+05:30', '2025-06-19 10:36:03.412739+05:30', NULL, 'Ahmedabad', 'Pune'),
(63, 5, 10, 2, '2024-06-26 09:00:00', '2024-06-26 12:30:00', 1106, 'Scheduled', 276.6, 459.0, 621.63, 732.15, 80, 84, 26, 12, 80, 84, 26, 12, '2025-06-19 10:36:03.412739+05:30', '2025-06-19 10:36:03.412739+05:30', NULL, 'Jaipur', 'Mumbai'),
(64, 4, 9, 4, '2024-06-29 09:00:00', '2024-06-29 12:00:00', 754, 'Scheduled', 152.37, 340.9, 462.02, 614.08, 88, 82, 35, 17, 88, 82, 35, 17, '2025-06-19 10:36:03.412739+05:30', '2025-06-19 10:36:03.412739+05:30', NULL, 'Pune', 'Chennai'),
(65, 1, 6, 4, '2024-06-28 10:00:00', '2024-06-28 13:30:00', 1285, 'Scheduled', 279.18, 453.73, 618.54, 754.03, 121, 52, 24, 8, 121, 52, 24, 8, '2025-06-19 10:36:03.412739+05:30', '2025-06-19 10:36:03.412739+05:30', NULL, 'Hyderabad', 'Chennai'),
(66, 1, 9, 3, '2024-06-27 06:00:00', '2024-06-27 10:00:00', 326, 'Scheduled', 184.56, 357.95, 475.94, 661.75, 104, 32, 32, 15, 104, 32, 32, 15, '2025-06-19 10:36:03.412739+05:30', '2025-06-19 10:36:03.412739+05:30', NULL, 'Pune', 'Bangalore'),
(67, 5, 10, 8, '2024-06-26 08:00:00', '2024-06-26 10:00:00', 1590, 'Scheduled', 254.77, 440.59, 573.22, 728.45, 129, 57, 16, 19, 129, 57, 16, 19, '2025-06-19 10:36:03.412739+05:30', '2025-06-19 10:36:03.412739+05:30', NULL, 'Jaipur', 'Kochi'),
(68, 3, 9, 5, '2024-06-28 08:00:00', '2024-06-28 12:00:00', 110, 'Scheduled', 156.96, 316.62, 511.6, 625.16, 103, 41, 32, 14, 103, 41, 32, 14, '2025-06-19 10:36:03.412739+05:30', '2025-06-19 10:36:03.412739+05:30', NULL, 'Pune', 'Kolkata'),
(69, 5, 2, 1, '2024-06-27 07:00:00', '2024-06-27 11:30:00', 561, 'Scheduled', 352.93, 483.69, 674.31, 786.54, 166, 84, 29, 12, 166, 84, 29, 12, '2025-06-19 10:36:03.412739+05:30', '2025-06-19 10:36:03.412739+05:30', NULL, 'Mumbai', 'Delhi'),
(70, 2, 7, 1, '2024-06-26 10:00:00', '2024-06-26 12:00:00', 630, 'Scheduled', 494.01, 671.77, 868.15, 1021.31, 169, 84, 35, 14, 169, 84, 35, 14, '2025-06-19 10:36:03.412739+05:30', '2025-06-19 10:36:03.412739+05:30', NULL, 'Ahmedabad', 'Delhi'),
(71, 3, 6, 3, '2024-06-28 08:00:00', '2024-06-28 10:30:00', 1462, 'Scheduled', 294.26, 475.69, 661.6, 854.49, 104, 48, 35, 13, 104, 48, 35, 13, '2025-06-19 10:36:03.412739+05:30', '2025-06-19 10:36:03.412739+05:30', NULL, 'Hyderabad', 'Bangalore'),
(72, 4, 8, 6, '2024-06-26 06:00:00', '2024-06-26 10:00:00', 1028, 'Scheduled', 242.16, 367.51, 470.77, 635.24, 177, 54, 28, 14, 177, 54, 28, 14, '2025-06-19 10:36:03.412739+05:30', '2025-06-19 10:36:03.412739+05:30', NULL, 'Kochi', 'Hyderabad'),
(73, 3, 4, 2, '2024-06-28 09:00:00', '2024-06-28 11:00:00', 1200, 'Scheduled', 296.4, 404.6, 541.47, 663.86, 156, 60, 26, 12, 156, 60, 26, 12, '2025-06-19 10:36:03.412739+05:30', '2025-06-19 10:36:03.412739+05:30', NULL, 'Chennai', 'Mumbai'),
(74, 3, 6, 4, '2024-06-26 07:00:00', '2024-06-26 11:30:00', 1697, 'Scheduled', 342.33, 542.33, 663.86, 763.95, 163, 87, 40, 11, 163, 87, 40, 11, '2025-06-19 10:36:03.412739+05:30', '2025-06-19 10:36:03.412739+05:30', NULL, 'Hyderabad', 'Chennai'),
(75, 1, 4, 9, '2024-06-26 08:00:00', '2024-06-26 12:30:00', 574, 'Scheduled', 343.28, 520.86, 698.8, 799.98, 159, 32, 29, 12, 159, 32, 29, 12, '2025-06-19 10:36:03.412739+05:30', '2025-06-19 10:36:03.412739+05:30', NULL, 'Chennai', 'Pune'),
(76, 1, 8, 10, '2024-06-27 09:00:00', '2024-06-27 11:30:00', 975, 'Scheduled', 412.55, 601.11, 775.07, 935.48, 165, 48, 13, 13, 165, 48, 13, 13, '2025-06-19 10:36:03.412739+05:30', '2025-06-19 10:36:03.412739+05:30', NULL, 'Kochi', 'Jaipur'),
(77, 2, 2, 4, '2024-06-27 09:00:00', '2024-06-27 11:30:00', 925, 'Scheduled', 270.63, 464.84, 605.39, 725.12, 88, 67, 27, 16, 88, 67, 27, 16, '2025-06-19 10:36:03.412739+05:30', '2025-06-19 10:36:03.412739+05:30', NULL, 'Mumbai', 'Chennai'),
(78, 4, 2, 1, '2024-06-26 07:00:00', '2024-06-26 11:00:00', 1257, 'Scheduled', 274.51, 434.49, 626.2, 805.29, 161, 75, 24, 6, 161, 75, 24, 6, '2025-06-19 10:36:03.412739+05:30', '2025-06-19 10:36:03.412739+05:30', NULL, 'Mumbai', 'Delhi'),
(79, 2, 1, 9, '2024-06-29 08:00:00', '2024-06-29 11:30:00', 1935, 'Scheduled', 404.15, 567.82, 756.22, 944.81, 87, 65, 13, 11, 87, 65, 13, 11, '2025-06-19 10:36:03.412739+05:30', '2025-06-19 10:36:03.412739+05:30', NULL, 'Delhi', 'Pune'),
(80, 2, 1, 2, '2024-06-28 09:00:00', '2024-06-28 11:00:00', 1219, 'Scheduled', 420.99, 528.96, 724.81, 912.57, 88, 75, 39, 13, 88, 75, 39, 13, '2025-06-19 10:36:03.412739+05:30', '2025-06-19 10:36:03.412739+05:30', NULL, 'Delhi', 'Mumbai'),
(81, 2, 7, 6, '2024-06-28 09:00:00', '2024-06-28 12:00:00', 241, 'Scheduled', 238.03, 416.8, 526.68, 704.52, 126, 42, 16, 12, 126, 42, 16, 12, '2025-06-19 10:36:03.412739+05:30', '2025-06-19 10:36:03.412739+05:30', NULL, 'Ahmedabad', 'Hyderabad'),
(82, 4, 6, 2, '2024-06-27 06:00:00', '2024-06-27 10:30:00', 205, 'Scheduled', 337.73, 484.57, 667.85, 780.53, 147, 85, 10, 8, 147, 85, 10, 8, '2025-06-19 10:36:03.412739+05:30', '2025-06-19 10:36:03.412739+05:30', NULL, 'Hyderabad', 'Mumbai'),
(83, 4, 2, 1, '2024-06-28 06:00:00', '2024-06-28 08:00:00', 659, 'Scheduled', 269.29, 423.87, 548.39, 736.58, 167, 62, 30, 18, 167, 62, 30, 18, '2025-06-19 10:36:03.412739+05:30', '2025-06-19 10:36:03.412739+05:30', NULL, 'Mumbai', 'Delhi'),
(84, 2, 1, 8, '2024-06-29 09:00:00', '2024-06-29 11:00:00', 1582, 'Scheduled', 358.42, 555.91, 689.05, 882.24, 175, 35, 31, 12, 175, 35, 31, 12, '2025-06-19 10:36:03.412739+05:30', '2025-06-19 10:36:03.412739+05:30', NULL, 'Delhi', 'Kochi'),
(85, 3, 2, 5, '2024-06-28 10:00:00', '2024-06-28 14:30:00', 1747, 'Scheduled', 396.72, 534.16, 640.99, 746.8, 169, 69, 20, 19, 169, 69, 20, 19, '2025-06-19 10:36:03.412739+05:30', '2025-06-19 10:36:03.412739+05:30', NULL, 'Mumbai', 'Kolkata'),
(86, 3, 1, 4, '2024-06-27 07:00:00', '2024-06-27 10:00:00', 1935, 'Scheduled', 224.77, 349.96, 499.63, 672.53, 92, 68, 30, 10, 92, 68, 30, 10, '2025-06-19 10:36:03.412739+05:30', '2025-06-19 10:36:03.412739+05:30', NULL, 'Delhi', 'Chennai'),
(87, 3, 2, 4, '2024-06-26 06:00:00', '2024-06-26 09:00:00', 136, 'Scheduled', 181.04, 283.94, 457.3, 598.69, 110, 33, 32, 19, 110, 33, 32, 19, '2025-06-19 10:36:03.412739+05:30', '2025-06-19 10:36:03.412739+05:30', NULL, 'Mumbai', 'Chennai'),
(88, 2, 7, 8, '2024-06-29 08:00:00', '2024-06-29 11:00:00', 772, 'Scheduled', 318.05, 426.58, 546.18, 673.76, 128, 47, 17, 5, 128, 47, 17, 5, '2025-06-19 10:36:03.412739+05:30', '2025-06-19 10:36:03.412739+05:30', NULL, 'Ahmedabad', 'Kochi'),
(89, 2, 9, 4, '2024-06-27 07:00:00', '2024-06-27 10:30:00', 1912, 'Scheduled', 324.8, 430.25, 603.39, 751.84, 173, 64, 18, 20, 173, 64, 18, 20, '2025-06-19 10:36:03.412739+05:30', '2025-06-19 10:36:03.412739+05:30', NULL, 'Pune', 'Chennai'),
(90, 3, 6, 2, '2024-06-27 10:00:00', '2024-06-27 14:00:00', 101, 'Scheduled', 393.61, 542.87, 649.95, 808.02, 90, 35, 17, 14, 90, 35, 17, 14, '2025-06-19 10:36:03.412739+05:30', '2025-06-19 10:36:03.412739+05:30', NULL, 'Hyderabad', 'Mumbai'),
(91, 2, 1, 10, '2024-06-28 08:00:00', '2024-06-28 10:00:00', 1684, 'Scheduled', 324.63, 457.86, 617.44, 749.95, 100, 59, 24, 20, 100, 59, 24, 20, '2025-06-19 10:36:03.412739+05:30', '2025-06-19 10:36:03.412739+05:30', NULL, 'Delhi', 'Jaipur'),
(92, 4, 4, 1, '2024-06-26 09:00:00', '2024-06-26 11:30:00', 1430, 'Scheduled', 467.33, 665.22, 770.35, 943.05, 108, 46, 18, 10, 108, 46, 18, 10, '2025-06-19 10:36:03.412739+05:30', '2025-06-19 10:36:03.412739+05:30', NULL, 'Chennai', 'Delhi'),
(93, 3, 9, 4, '2024-06-29 07:00:00', '2024-06-29 10:00:00', 106, 'Scheduled', 468.18, 619.15, 746.84, 929.77, 93, 43, 12, 5, 93, 43, 12, 5, '2025-06-19 10:36:03.412739+05:30', '2025-06-19 10:36:03.412739+05:30', NULL, 'Pune', 'Chennai'),
(94, 5, 10, 4, '2024-06-28 09:00:00', '2024-06-28 11:00:00', 1186, 'Scheduled', 290.3, 458.9, 640.8, 801.03, 146, 55, 11, 14, 146, 55, 11, 14, '2025-06-19 10:36:03.412739+05:30', '2025-06-19 10:36:03.412739+05:30', NULL, 'Jaipur', 'Chennai'),
(95, 2, 5, 4, '2024-06-28 09:00:00', '2024-06-28 12:00:00', 641, 'Scheduled', 458.2, 636.05, 831.55, 1016.34, 159, 58, 35, 10, 159, 58, 35, 10, '2025-06-19 10:36:03.412739+05:30', '2025-06-19 10:36:03.412739+05:30', NULL, 'Kolkata', 'Chennai'),
(96, 3, 5, 7, '2024-06-29 06:00:00', '2024-06-29 08:00:00', 1207, 'Scheduled', 443.78, 617.31, 750.97, 910.98, 175, 80, 29, 7, 175, 80, 29, 7, '2025-06-19 10:36:03.412739+05:30', '2025-06-19 10:36:03.412739+05:30', NULL, 'Kolkata', 'Ahmedabad'),
(97, 5, 10, 2, '2024-06-29 07:00:00', '2024-06-29 10:00:00', 1418, 'Scheduled', 358.23, 461.11, 616.24, 785.2, 147, 30, 36, 7, 147, 30, 36, 7, '2025-06-19 10:36:03.412739+05:30', '2025-06-19 10:36:03.412739+05:30', NULL, 'Jaipur', 'Mumbai'),
(98, 1, 7, 1, '2024-06-26 09:00:00', '2024-06-26 13:00:00', 1238, 'Scheduled', 403.75, 521.79, 621.9, 792.38, 107, 38, 13, 5, 107, 38, 13, 5, '2025-06-19 10:36:03.412739+05:30', '2025-06-19 10:36:03.412739+05:30', NULL, 'Ahmedabad', 'Delhi'),
(99, 3, 8, 10, '2024-06-29 10:00:00', '2024-06-29 12:00:00', 1056, 'Scheduled', 341.67, 483.97, 672.06, 776.06, 103, 65, 17, 17, 103, 65, 17, 17, '2025-06-19 10:36:03.412739+05:30', '2025-06-19 10:36:03.412739+05:30', NULL, 'Kochi', 'Jaipur');


INSERT INTO public.travellers (id, booking_id, title, first_name, last_name, dob, nationality) VALUES
(1, 15, 'Mr', 'Chirag', 'Goyal', '2000-10-11 00:00:00+05:30', 'India'),
(2, 15, 'Mrs', 'Devansh', 'president)', '2005-10-25 00:00:00+05:30', 'Other'),
(3, 15, 'Ms', 'Devansh', 'president', '2005-10-25 00:00:00+05:30', 'Other');


INSERT INTO public.users (id, created_at, updated_at, deleted_at, email, password, name, is_admin) VALUES
(2, '2025-06-17 16:29:01.889882+05:30', '2025-06-17 16:29:01.889882+05:30', NULL, 'goyalchirag573@gmail.com', '$2a$10$UxZqFZUCZrhkWwduur0bruC7G7ZlDEAQC2sRwLGogepnhJziEXjzu', 'Chirag Goyal', false),
(1, '2025-06-17 16:28:49.805774+05:30', '2025-06-17 16:28:49.805774+05:30', NULL, 'dangsterplayer0010@gmail.com', '$2a$10$v9xoIDDwp76UQZDW4EVlo.d7nIZj8Wa53GkS468Ht8.g2sYIcu3mK', 'Chirag', true);


SELECT pg_catalog.setval('public.aircrafts_aircraft_id_seq', 5, true);
SELECT pg_catalog.setval('public.airports_airport_id_seq', 10, true);
SELECT pg_catalog.setval('public.bookings_id_seq', 15, true);
SELECT pg_catalog.setval('public.flights_id_seq', 156, true);
SELECT pg_catalog.setval('public.travellers_id_seq', 3, true);
SELECT pg_catalog.setval('public.users_id_seq', 2, true);

ALTER TABLE ONLY public.aircrafts
    ADD CONSTRAINT aircrafts_pkey PRIMARY KEY (aircraft_id);

ALTER TABLE ONLY public.airports
    ADD CONSTRAINT airports_pkey PRIMARY KEY (airport_id);

ALTER TABLE ONLY public.bookings
    ADD CONSTRAINT bookings_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.flights
    ADD CONSTRAINT flights_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.travellers
    ADD CONSTRAINT travellers_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_email_key UNIQUE (email);

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);

CREATE INDEX idx_bookings_deleted_at ON public.bookings USING btree (deleted_at);
CREATE INDEX idx_flights_deleted_at ON public.flights USING btree (deleted_at);
CREATE INDEX idx_users_deleted_at ON public.users USING btree (deleted_at);

ALTER TABLE ONLY public.travellers
    ADD CONSTRAINT fk_bookings_travellers FOREIGN KEY (booking_id) REFERENCES public.bookings(id) ON DELETE CASCADE;

ALTER TABLE ONLY public.bookings
    ADD CONSTRAINT fk_flights_bookings FOREIGN KEY (flight_id) REFERENCES public.flights(id);

ALTER TABLE ONLY public.bookings
    ADD CONSTRAINT fk_users_bookings FOREIGN KEY (user_id) REFERENCES public.users(id);

ALTER TABLE ONLY public.flights
    ADD CONSTRAINT flights_aircraft_id_fkey FOREIGN KEY (aircraft_id) REFERENCES public.aircrafts(aircraft_id);

ALTER TABLE ONLY public.flights
    ADD CONSTRAINT flights_arrival_airport_id_fkey FOREIGN KEY (arrival_airport_id) REFERENCES public.airports(airport_id);

ALTER TABLE ONLY public.flights
    ADD CONSTRAINT flights_departure_airport_id_fkey FOREIGN KEY (departure_airport_id) REFERENCES public.airports(airport_id);

