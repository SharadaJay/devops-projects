## Steps and Procedure

1. Clone the repository.
    ```bash
           git clone -b project https://course-gitlab.tuni.fi/compse140-fall2023/fnshja.git
    ```

2. Change the directory.
    ```bash
           cd fnshja
    ```
   
3. Build the system.
    ```bash
           docker-compose build --no-cache
    ```

4. Run the system.
    ```bash
           docker-compose up -d
    ```
   
5. Wait for a bit until all the services become ready (approximately 25 seconds) and then use curl/Postman to test the system (See **EndReport.pdf** for more details).