package main

const indexHTML = `
<!DOCTYPE html>
<html>
<head>
    <title>Servidor Web en Go</title>
    <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/4.5.2/css/bootstrap.min.css">
    <style>
        .fondo-cielo {
            background-color: #87CEEB; 
        }
        .card {
            background-color: #000; 
            color: #fff;
            margin: 10px; 
        }
        .card-img-top {
            margin: -10px -10px 0 -10px; 
        }
        .card-title {
            color: #fff; 
            margin: 0; 
            padding: 10px; 
        }
        .recuadro-inferior {
            background-color: #000; 
            color: #fff; 
            text-align: center; 
            padding: 10px; 
            position: bottom; 
            bottom: 0; 
        }
    </style>
</head>
<body>
    <div class="container">

        <div class="fondo-cielo text-white p-3">
            <h1 >Bienvenido al Servidor de Imágenes en Go!!</h1>
        </div>

        <div class="text-center font-weight-bold mt-4">
            Nombre de host: {{ .Hostname }}
        </div>

        <br/>
        <div class="row">
            {{ range .Imagenes }}
            <div class="col-md-4 mb-4">
                <div class="card">
                    <img class="card-img-top" src="data:image/png;base64,{{ .Base64 }}" alt="{{ .Nombre }}">
                    <div class="card-body">
                        <h5 class="card-title">{{ .Nombre }}</h5>
                    </div>
                </div>
            </div>
            {{ end }}
        </div>

        <!-- Recuadro inferior centrado y sin conectar con las paredes -->
        <div class="recuadro-inferior p3">
            <p>Computación en la nube</p>
            <p>Andrés Mauricio Dussán Bastidas</p>
            <p>Octubre de 2023</p>
        </div>
        
    </div>
</body>
</html>


`
