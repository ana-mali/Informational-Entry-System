## Installation

Latest version of GoLang (1.20+): 
[Download and install](https://go.dev/doc/install)

## Usage
There are multiple ways to use the Notes Application, please choose the one you desire. Make sure you are located in the NotesApp folder.

For the standard CLI application, please run the following to get an understanding of how to run the program. 
```powershell
go run . app help
```
For the API serviced version, the routing will be slightly different with improved and additional features. Run the following to understand how to use the enhanced application. M
```powershell
go run ./cmd/cli app help
```
## Command Structure Overview
The CLI follows this general structure:
```powershell
app <resource> <subresource> <command> <args> [flags]
```
Where the resource can be notes, tasks or lists, the sub resource is usually associated with items connected to the list resource, the command can be an add, remove or edit and the rest are the parameters used to perform the action which are specific to the resource and sub resource. 
If your string argument has spaces, use quotations " " so that the command line can recognize it as part of one argument. 

## Resources and Commands
Notes: Single editable field. Text field is mandatory.
```powershell
app notes edit <noteID> --text "Update a note text"
app notes add "Make a list"
app notes remove <noteID>
app notes list
```
Tasks: Entries that can be tracked for completion time and priority levels. Setting priority and due date fields are completely optional.
```powershell
app tasks edit <taskID> --priority high --due YYYY-MM-DD
app tasks edit <taskID> --clear priority --clear due
app tasks add "Call Dentist" --priority medium 
app tasks remove <taskID>
app tasks list
```
Lists: Holds a list of items. The only field that can be edited is the name, and this must always exist in a list structure. 
```powershell
app lists edit <listID> --name "new name for list"
app lists add "To Do List"
app lists remove <listID>
app lists list
```
Items: Belong to a specific list. Can be 'checked' as a way for the user to tick off items from their lists. Items can be seen through using the list command for the lists resource. 
```powershell 
app lists item edit <listID> <itemID> --text "Shoes" --check true
app lists item add <listID> "Buy Shoes" 
app lists item remove <listID> <itemID> 

```

## Future Updates

With additional resources available, this application can be expanded in a variety of ways. Such as incorporating even more resources such as reminder popup functionalities and similar calendar features. Furthermore, making this application more user friendly through a designated front end or improved functionalities can make this go from a simple CLI application to a web platform for personal information management. 
