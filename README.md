# Strava2GSheets 🏃🏼‍♀️ ➡️ 📊

**Strava2GSheets** is what happens when a developer spends way too much time 
automating a task that barely takes a minute to do manually.

Instead of spending 60 seconds a day updating my training diary, I spent hours 
(okay, weeks) writing this Go script to do it for me.

Because why do something manually when you can automate it and feel ridiculously
proud every time it runs?

### Could I have just updated my Google Sheet manually?
*Sure.*

### Would that have taken less time than writing this script?
*Absolutely.*

## How It Works

### Strava API:

- I use Strava’s API to pull my workouts.
- Setting it up involves creating an app on the Strava Developer Portal and grabbing the client ID and secret.

### Google Sheets API:

- I enabled the Sheets API via the Google Cloud Console.
- Downloaded credentials to give the script access to my spreadsheet.

### GitHub Workflow:

- I’ve added a scheduled GitHub Actions workflow to run this script automatically.
- The workflow triggers daily (or as scheduled) and can also be run manually if needed.
