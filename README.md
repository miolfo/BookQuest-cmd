# finna-goodreads
Integration between finna and goodreads to check what books are available in a specific location. Can be used to read an exported .csv file from goodreads,
and then checking if they can be found from a local library using Finna public api. The tool only checks books in your want-to-read -shelf.

Usage

`
go run . ./resources/goodreads_library_export.csv 0/Helmet/ ./resources/result.csv
`

First parameter is the name of the input file. Second parameter is the building to look for (for example, 0/Helmet/ includes all Helsinki city libraries).
Third parameter is the output file.
Output generates a row like 

`
"Children of Dune (Dune, #3)",Frank Herbert,https://www.finna.fi/Record/helmet.2404461
`

when the book was found, and if not, the finna url will be omitted. 