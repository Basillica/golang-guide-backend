# Golang Tutorial

### GoLang Fundamentals (fundamentals)

-   GoLang for Backend Development: Advantages and Use Cases
-   Setting Up Your GoLang Development Environment
-   Core GoLang Concepts: Variables, Data Types, Operators, Control Flow, Functions
-   Important concepts
    -- Concurrency
    -- Channels
    -- Async processes

### Building Blocks for Web Applications (webapp)

-   Introduction to Web Development with GoLang: HTTP Requests & Responses
    -- Development with Gin framework
-   Routing: Defining Endpoints and Handling Requests
-   Working with Databases: Choosing and Using a Database (SQL vs. NoSQL)
-   Connecting to your database
    -- Working with ORMs
    -- Writting raw sql queries

### Data Management and Security (security)

-   Building CRUD Operations: Create, Read, Update, Delete
-   User Authentication and Authorization: Securing Your Application
    -- Password Hashing (securely storing passwords)
    -- Error Handling: Gracefully Handling Errors in Web Applications

### Advanced Web Development Techniques (techniques)

-   Templates: Building Dynamic User Interfaces
-   Working with Data Formats: JSON and Beyond (handling data exchange)
-   Session Management (maintaining user state)
-   cookies

### Heirarchical RBAC implementation in Golang (h-rbac)

### User Interaction and Communication (comms)

-   Forms: Processing User Input and Handling Submissions
-   WebSockets: Enabling Real-time Communication

### GoLang Web Development Ecosystem (ecosystem)

-   Dependency Management with Go Modules (managing external libraries)

### Building Robust Web Applications (robust)

-   Middleware: Interceptors for Request Processing (common tasks and custom functionality)
-   Testing Your Backend Web Application: Unit and Integration Tests
    -- Mocking Dependencies for Effective Testing

### Scaling and Performance Optimization (optimization)

-   Caching Strategies for Improved Response Times
    -   In-memory
    -   Distributed
    -   File-based
    -   CDN
-   Logging and Monitoring: Tracking Application Behavior
    -   Stucuted logging with Zap
    -   App metrics with Prometheus
    -   Distrubuted tracing with Jaeger
    -   Log aggregation with ELK stack
    -   Alerting and notification with Prometheus Alertmanager

### Advanced Security Practices (advanced-practices)

-   Input Validation: Sanitizing User Input to Prevent Attacks

    -   SQL injection
    -   Cross-site scripting XSS
    -   Remote code execution RCE
    -   Path traversal
    -   Command injection

    Solution:

    1. User input sanitization and validation
       a. Whitelisting and Blacklisting
       b. Regular expressions
       c. Third-party libraries
    2. Context-aware prevention
       a. SQL injection prevention
       b. XSS prevention
       c. RCE prevention

-   SQL Injection Prevention: Protecting Your Database

### Deployment Strategies (strategies)

-   Cloud Deployment Platforms (Heroku, AWS, Google Cloud)
-   Containerization with Docker (packaging your application for deployment)
