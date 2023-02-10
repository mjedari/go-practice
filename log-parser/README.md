## Log Parser

This app is a simple app to parse large log file *(for instance in 100 millions lines)* and persist them in database.
The important part would be we should read lines by chunk and store them in chunk async. 

The architecture we use is clean architecture.