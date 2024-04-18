# youtubeApiFetch

End point ```/videosPaginated/1``` accepts the page number and returns top 5 result based on latest released time

End point ```/getVideo?description=Manu Carreño&title=CARREÑO tras la VICTORIA del MADRID ante el CITY: %26quot;Si lo hace el %26%2339;CHOLO%26%2339;%26quot;``` accepts description and title of the video

The Entire code runs on port 3000. 

Sample curl 
curl --location 'localhost:3000/getVideo?description=Manu%20Carre%C3%B1o&title=CARRE%C3%91O%20tras%20la%20VICTORIA%20del%20MADRID%20ante%20el%20CITY%3A%20%26quot%3BSi%20lo%20hace%20el%20%26%2339%3BCHOLO%26%2339%3B%26quot%3B'


curl --location 'localhost:3000/videosPaginated/1'