# devops-course
Git Repository for Continuous Development and Deployment - DevOps Master's Course at Tampere University

<br>

**project** branch holds the implementation of the instructions provided in the *project-instructions* PDF file. This is the final project of the course. Please refer to the instructions here [project-instructions](./project-instructions.pdf).

<br>



## Steps and Procedure

1. Clone the repository.
    ```bash
           git clone -b project https://github.com/SharadaJay/devops-course.git
    ```

2. Change the directory.
    ```bash
           cd devops-course
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