CREATE DATABASE IF NOT EXISTS eventtct;
USE eventtct;

CREATE TABLE IF NOT EXISTS eventsfintech (
    id INT AUTO_INCREMENT PRIMARY KEY,
    event_code VARCHAR(50) NOT NULL,
    event_amount DECIMAL(15,2) NOT NULL,
    event_quantity INT NOT NULL,
    site_id VARCHAR(10) NOT NULL,
    started_date DATE NOT NULL
);

INSERT INTO eventsfintech (event_code, event_amount, event_quantity, site_id, started_date) VALUES
('TC_CAR_001', 15000.00, 30, 'MLM', '2024-06-01'),
('TC_CAR_004', 5250.50, 28, 'MLM', '2024-06-01'),
('TC_COB_001', 35000.00, 28, 'MLB', '2024-06-02'),
('TC_COB_002', 27500.00, 180, 'MLB', '2024-06-02'),
('TC_VTO_003_transacciones_positivas', 18500.00, 180, 'MLM', '2024-06-03'),
('TC_VTO_003_transacciones_negativas', -7500.00, 180, 'MLM', '2024-06-03'); 