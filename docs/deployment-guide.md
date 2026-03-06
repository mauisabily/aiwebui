# Deployment and Usage Guide

## Overview

This guide provides instructions for deploying and using the AI WebUI application. The application consists of a Go backend server and a Vue.js frontend that connects to a local Ollama service.

## System Requirements

### Server Requirements
- Linux, macOS, or Windows operating system
- Go 1.21 or higher
- MySQL 8.0 or higher
- Access to Ollama service at http://192.168.1.50:11434

### Client Requirements
- Modern web browser (Chrome, Firefox, Safari, Edge)
- Internet connection (for CDN resources)

## Deployment Steps

### 1. Database Setup

#### Install MySQL
On Ubuntu/Debian:
```bash
sudo apt update
sudo apt install mysql-server
sudo mysql_secure_installation
```

On CentOS/RHEL:
```bash
sudo yum install mysql-server
sudo systemctl start mysqld
sudo mysql_secure_installation
```

#### Configure Database
Connect to MySQL as root:
```bash
mysql -u root -p
```

Create database and user:
```sql
CREATE DATABASE ai_kpst;
CREATE USER 'ai_kpst'@'localhost' IDENTIFIED BY 'c61762a01f19d8';
GRANT ALL PRIVILEGES ON ai_kpst.* TO 'ai_kpst'@'localhost';
FLUSH PRIVILEGES;
EXIT;
```

#### Run Database Migrations
Execute the SQL script from [MySQL Setup](mysql-setup.md):
```bash
mysql -u ai_kpst -p ai_kpst < migrations/001_create_tables.sql
```

### 2. Backend Deployment

#### Clone Repository
```bash
git clone <repository-url>
cd aiwebui
```

#### Configure Application
Create `configs/config.yaml`:
```yaml
server:
  port: 8080
  host: "0.0.0.0"

mysql:
  host: "localhost"
  port: 3306
  username: "ai_kpst"
  password: "c61762a01f19d8"
  database: "ai_kpst"
  charset: "utf8mb4"

ollama:
  base_url: "http://192.168.1.50:11434"
  default_model: "llama3"

rag:
  chunk_size: 1000
  chunk_overlap: 200
  max_results: 5
```

#### Build Backend
```bash
go mod tidy
go build -o aiwebui cmd/server/main.go
```

#### Run Backend
```bash
./aiwebui
```

The server will start on port 8080 by default.

### 3. Frontend Deployment

#### Install Dependencies
```bash
cd web/vue
npm install
```

#### Production Build
```bash
npm run build
```

This will generate static files in the `dist` directory.

#### Serve Frontend
You can serve the frontend files using any web server. For example, with Nginx:

1. Install Nginx:
```bash
sudo apt install nginx
```

2. Copy built files:
```bash
sudo cp -r dist/* /var/www/html/
```

3. Configure Nginx to proxy API requests to the Go backend:
```nginx
server {
    listen 80;
    server_name your-domain.com;
    
    location / {
        root /var/www/html;
        try_files $uri $uri/ /index.html;
    }
    
    location /api/ {
        proxy_pass http://localhost:8080/;
    }
}
```

### 4. Ollama Setup

Ensure Ollama is running and accessible at http://192.168.1.50:11434.

Pull a model if needed:
```bash
ollama pull llama3
```

## Usage Instructions

### First-time Setup

1. Open your web browser and navigate to the application URL
2. The application will automatically detect available models from Ollama
3. Start a new conversation by typing in the message input

### Chatting with AI

1. **Starting a Conversation**
   - Click "New Chat" in the sidebar
   - Type your message in the input box at the bottom
   - Press Enter or click the send button

2. **Switching Models**
   - Click the model selector in the header
   - Choose from available models
   - New messages will use the selected model

3. **Managing Conversations**
   - View past conversations in the sidebar
   - Click on any conversation to resume it
   - Delete conversations using the trash icon

### Knowledge Base Management

1. **Creating a Knowledge Base**
   - Click "Knowledge Bases" in the sidebar
   - Click "New Knowledge Base"
   - Enter name and description
   - Click "Create"

2. **Adding Documents**
   - Select a knowledge base
   - Click "Upload Document"
   - Choose a file from your computer
   - Click "Upload"

3. **Searching Knowledge Base**
   - Select a knowledge base
   - Use the search bar to find relevant documents
   - Results show document titles and relevance scores

### Enabling RAG

1. **Toggle RAG for Conversation**
   - Open a conversation
   - Toggle the "Enable RAG" switch
   - Select knowledge bases to use for retrieval

2. **Viewing Sources**
   - When RAG is enabled, responses will include source documents
   - Sources are displayed below each AI response
   - Click on sources to view document details

### Customizing Settings

1. **Accessing Settings**
   - Click the settings icon in the header
   - Adjust preferences in the modal

2. **Available Settings**
   - Theme (light/dark)
   - Default model selection
   - Auto-save conversations
   - Language preferences

## Troubleshooting

### Common Issues

1. **Cannot Connect to Database**
   - Check MySQL service status: `sudo systemctl status mysql`
   - Verify database credentials in config.yaml
   - Ensure MySQL user has proper permissions

2. **Ollama Connection Failed**
   - Verify Ollama is running: `systemctl status ollama`
   - Check Ollama URL in config.yaml
   - Ensure firewall allows access to Ollama port

3. **Frontend Not Loading**
   - Check web server configuration
   - Verify static files are in correct location
   - Check browser console for errors

4. **Slow Responses**
   - Check system resources (CPU, memory)
   - Verify Ollama model performance
   - Consider using smaller models for faster responses

### Logs and Monitoring

1. **Backend Logs**
   - Check terminal output where backend is running
   - Look for error messages or warnings

2. **Database Logs**
   - Check MySQL error logs: `/var/log/mysql/error.log`

3. **Frontend Logs**
   - Check browser developer console
   - Look for network errors or JavaScript exceptions

## Maintenance

### Updating the Application

1. Pull latest changes:
```bash
git pull origin main
```

2. Update dependencies:
```bash
go mod tidy
cd web/vue && npm install
```

3. Rebuild and restart:
```bash
go build -o aiwebui cmd/server/main.go
# Restart the service
```

### Backing Up Data

1. **Database Backup**
```bash
mysqldump -u ai_kpst -p ai_kpst > backup.sql
```

2. **Document Backup**
   - Copy document files from the storage location
   - Export knowledge base configurations

### Scaling Considerations

For production deployments with many users:

1. **Database Optimization**
   - Use connection pooling
   - Implement read replicas for heavy read workloads
   - Optimize indexes for search queries

2. **Application Scaling**
   - Run multiple backend instances behind a load balancer
   - Use caching for frequently accessed data
   - Implement rate limiting

3. **Storage Scaling**
   - Use cloud storage for document files
   - Implement CDN for static assets
   - Consider specialized vector databases for large embeddings

## Security Considerations

### Authentication
The current version does not include authentication. For production use:

1. Implement user authentication
2. Add role-based access control
3. Use HTTPS for all communications

### Data Protection
1. Encrypt sensitive data at rest
2. Use parameterized queries to prevent SQL injection
3. Validate and sanitize all user inputs
4. Regularly update dependencies

## Support

For issues not covered in this guide:

1. Check the GitHub issues page
2. Contact the development team
3. Refer to documentation in the `docs/` directory

This deployment and usage guide should provide all necessary information to successfully deploy and use the AI WebUI application.