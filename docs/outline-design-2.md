# Crawler System Design for GitHub PRs Labeled as "Bugs"

## 1. Introduction
This document outlines the design for a crawler system that retrieves Pull Requests (PRs) labeled as "bugs" from GitHub, leveraging the GitHub REST API with a personal access token for authentication, storing data in a PostgreSQL database, and initiating the crawl process through a main function without a scheduler.

## 2. System Overview
The system is designed to query PRs using GitHub's REST API, focusing specifically on those tagged with the "bug" label. Key components of the system include API integration for data retrieval, PostgreSQL for data storage, and rate limit management using the `X-RateLimit-Remaining` HTTP header provided by GitHub.

## 3. Design Components

### 3.1 GitHub API Integration
- **Objective:** Fetch PRs labeled as "bugs" using GitHub's REST API.
- **Authentication:** Utilize a personal access token for authenticated API requests to increase the rate limit allowance.
- **Implementation:**
  - Construct API requests to search for PRs with the "bug" label within the targeted repositories.
  - Handle pagination to ensure complete data retrieval.
  - Parse the `X-RateLimit-Remaining` header to monitor and respect the API's rate limit.

### 3.2 Data Storage (PostgreSQL)
- **Objective:** Efficiently store and manage the retrieved PR data.
- **Implementation:**
  - Design a relational schema suitable for PR data, including fields such as PR ID, title, author, creation date, and status.
  - Implement SQL queries for inserting new records and updating existing ones to keep the dataset current.

### 3.3 Crawler Execution
- **Objective:** Provide a mechanism to start the crawler manually.
- **Implementation:**
  - Develop a main function that initiates the crawling process when executed.
  - Ensure that the function can perform a complete crawl cycle, fetching all relevant PRs and updating the database accordingly.

### 3.4 Rate Limit Management
- **Objective:** Ensure the crawler respects GitHub's API rate limits.
- **Implementation:**
  - Use the `X-RateLimit-Remaining` HTTP header to track the remaining number of requests allowed.
  - Implement logic to pause or delay requests when approaching the rate limit, resuming only when the limit is reset.

## 4. Additional Considerations

### 4.1 Error Handling and Data Integrity
- Implement error handling for API request failures, data parsing errors, and database interactions to ensure the crawler's robustness.
- Validate incoming data for completeness and accuracy before insertion into the database.

### 4.2 Monitoring and Logging
- Integrate logging to capture the crawler's operational details, including start and end times, number of PRs fetched, and any encountered errors.
- Use logging data to monitor crawler performance and troubleshoot issues as needed.

### 4.3 Security and Compliance
- Securely store the personal access token and sensitive database credentials, ensuring they are not exposed in the codebase or logs.
- Adhere to GitHub's terms of service and data protection regulations in the handling and storage of PR data.

### 4.4 Testing and Deployment
- Perform unit and integration testing to validate each component of the crawler, ensuring its functionality and reliability.
- Prepare for deployment, considering the execution environment and any dependencies required by the crawler.

### 4.5 Documentation
- Document the system architecture, codebase, and operational procedures to facilitate understanding and maintenance of the crawler.

## 5. Conclusion
This revised outline provides a focused approach for developing a GitHub PR crawler that respects specified constraints, including authentication, data storage, execution method, and rate limit management. The design ensures efficient data retrieval and storage while maintaining compliance with GitHub's API usage policies.
