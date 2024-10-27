# Database Setup Guide

## Database Location
The application initializes the database from the user's configuration folder. The specific location varies depending on the operating system:

- Linux: `~/.config/automathaus/automathaus_data/`
- macOS: `~/Library/Application Support/automathaus/automathaus_data/`
- Windows: `C:\Users\<YourUsername>\AppData\Roaming\automathaus\automathaus_data\`

## Initial Setup
1. Ensure you have the necessary permissions to access and modify the user configuration folder.
2. On first run, the application will automatically create the database in the appropriate location. However, to set up the initial structure and data:
   a. Locate the database files in this folder (db/).
   b. Copy these files to the `/automathaus_data` folder in your user configuration directory.
   c. The application will then use these files to initialize the database with the necessary tables and initial data.

Note: If you skip this step, the database will be created but will be empty, lacking the required structure for the application to function properly.



## Backup and Restore
It's recommended to regularly backup your database:

1. Stop the application.
2. Copy the entire `automathaus_data` folder to a secure location.
3. To restore, replace the existing `automathaus_data` folder with your backup.

## Admin user credentials
The application comes with default admin credentials. For security reasons, it is strongly recommended that you change these credentials immediately after your first login:

- Username: admin@automathaus.dev
- Password: AutomatPass1234

Remember to choose a unique, complex password and store it securely. Never share your admin credentials with others.


## Troubleshooting
If you encounter database-related issues:

1. Check the application logs for any error messages.
2. Verify the database file permissions.
3. Ensure sufficient disk space in the user configuration folder.

For further assistance, please contact the development team.