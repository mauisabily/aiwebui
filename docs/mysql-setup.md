# MySQL Database Setup

## Overview
This document describes how to set up the MySQL database for the AI WebUI application.

## Prerequisites
- MySQL server 8.0 or higher
- Database credentials:
  - Host: localhost
  - Port: 3306
  - Username: ai_kpst
  - Password: c61762a01f19d8
  - Database: ai_kpst

## Database Initialization

Run the following SQL script to create the necessary tables and indexes:

```sql
-- Create users table
CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(255) UNIQUE NOT NULL,
    email VARCHAR(255) UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- Create user_settings table
CREATE TABLE user_settings (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT NOT NULL,
    settings JSON,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

-- Create conversations table
CREATE TABLE conversations (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT NOT NULL,
    title VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

-- Create messages table
CREATE TABLE messages (
    id INT AUTO_INCREMENT PRIMARY KEY,
    conversation_id INT NOT NULL,
    role ENUM('user', 'assistant', 'system') NOT NULL,
    content TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (conversation_id) REFERENCES conversations(id) ON DELETE CASCADE
);

-- Create knowledge_bases table
CREATE TABLE knowledge_bases (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- Create documents table
CREATE TABLE documents (
    id INT AUTO_INCREMENT PRIMARY KEY,
    knowledge_base_id INT NOT NULL,
    title VARCHAR(255) NOT NULL,
    content LONGTEXT,
    source_type VARCHAR(50),
    source_url TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (knowledge_base_id) REFERENCES knowledge_bases(id) ON DELETE CASCADE
);

-- Create document_chunks table
CREATE TABLE document_chunks (
    id INT AUTO_INCREMENT PRIMARY KEY,
    document_id INT NOT NULL,
    content TEXT NOT NULL,
    chunk_order INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (document_id) REFERENCES documents(id) ON DELETE CASCADE
);

-- Create chunk_embeddings table
CREATE TABLE chunk_embeddings (
    id INT AUTO_INCREMENT PRIMARY KEY,
    chunk_id INT NOT NULL,
    embedding BLOB, -- Vector representation
    model_used VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (chunk_id) REFERENCES document_chunks(id) ON DELETE CASCADE
);

-- Create conversation_context table
CREATE TABLE conversation_context (
    id INT AUTO_INCREMENT PRIMARY KEY,
    conversation_id INT NOT NULL,
    context_type VARCHAR(100) NOT NULL,
    context_data TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (conversation_id) REFERENCES conversations(id) ON DELETE CASCADE
);

-- Create indexes for better performance
CREATE INDEX idx_conversations_user_id ON conversations(user_id);
CREATE INDEX idx_messages_conversation_id ON messages(conversation_id);
CREATE INDEX idx_messages_created_at ON messages(created_at);
CREATE INDEX idx_documents_knowledge_base_id ON documents(knowledge_base_id);
CREATE INDEX idx_document_chunks_document_id ON document_chunks(document_id);

-- Full-text search indexes
CREATE FULLTEXT INDEX idx_documents_content ON documents(content);
CREATE FULLTEXT INDEX idx_document_chunks_content ON document_chunks(content);

-- Insert default knowledge base
INSERT INTO knowledge_bases (name, description) VALUES ('Default Knowledge Base', 'Main knowledge base for general information');
```

## Connection Configuration

The application will connect to MySQL using the following configuration:

```yaml
mysql:
  host: localhost
  port: 3306
  username: ai_kpst
  password: c61762a01f19d8
  database: ai_kpst
  charset: utf8mb4
```

## Notes

1. Make sure the MySQL user has sufficient privileges to create tables and indexes.
2. The `embedding` column in the `chunk_embeddings` table is defined as BLOB to store vector representations. Depending on your MySQL version and requirements, you might need to adjust this.
3. Full-text search indexes are created on content columns to enable efficient keyword-based search functionality.
4. Foreign key constraints with cascade deletion ensure data integrity when parent records are removed.