CREATE TABLE IF NOT EXISTS settings (
    `key` VARCHAR(255) PRIMARY KEY,
    `value` TEXT NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_VALUE ON UPDATE CURRENT_VALUE
);

INSERT IGNORE INTO settings (`key`, `value`) VALUES 
('llm_mode', 'ollama'),
('ollama_url', 'http://192.168.1.50:11434'),
('airllm_url', 'http://localhost:8000'),
('default_model', 'llama3');
