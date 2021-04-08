# Golang - Time Zone API  
  
A simple webservice writen in Golang that interacts with the _Time_ and _Location_ packages.
Functionality is then exposed API based.  
  
## Example  

Get the current time of a specific UTC zone:
> get /timezone/{zone}

Get the time of your current time zone:
> get /timezone

Get a list of all selectable time zones:
> get /timeszones

