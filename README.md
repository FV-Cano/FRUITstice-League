<h1 align="center">🍓<span color="#bf2136">FRUIT</span>stice League </h1>

<h2>Descripción</h2>
<p>¡Bienvenidos al sistema oficial de gestión de inventario de frutas de La Liga de la Justicia!</p>
<p>Este sistema cumple con las funcionalidades de control de inventario, registro de ventas y generación de reportes de las Superfrutas de la Liga.</p>

<h2>Instrucciones de Ejecución</h2>
<ol>
  <li>Descargue el código en un archivo .zip desde GitHub presionando el botón verde "Code" y seleccionando la opción de 'Download Zip'.</li>
  <br>
  <li>Para ejecutar el programa asegúrese de contar con un sistema operativo Windows:</li>
  <br>
  <ol>
    <li>Descomprima el .zip descargado.</li>
    <br>
    <li>Ingrese a la carpeta <strong>'app'</strong>.</li>
    <br>
    <li>Ejecute el archivo <strong>'main.exe'</strong>.</li>
  </ol>
  <br>
  <li>Al ejecutar el programa este le dará la opción de realizar una carga automática de inventario o una carga manual. Si elige una carga automática, el inventario se cargará con los siguientes valores:</li>
  <br>
    
  ```bash
  batnana,20,19.65,2024-07-14
  straw-Bat-Berry,15,12.50,2024-06-30
  supergrape,25,8.99,2024-06-25
  aquamelon,12,15.25,2024-07-05
  green kiwi-lantern,22,9.45,2024-07-12
  cybapple,17,14.20,2024-06-28
  plum woman,14,16.30,2024-07-15
  ```
  <br>
  Si elige una carga manual podrá realizar una carga inicial de inventario con los productos que desee siguiendo el formato:
  <br>

  ```bash
  NombreProducto, Stock, PrecioUnitario, FechaVto
  ```
  <br>
  <li>Tras la carga inicial del inventario se desplegará un menú de opciones donde podrá explorar el resto de funcionalidades del sistema.</li>
</ol>

<h2>Observaciones</h2>
<p>El sistema cuenta con una API RESTful con 3 endpoints para realizar distintas funcionalidades. Si bien se puede acceder a estas funcionalidades mediante el uso del sistema, también tiene la opción de utilizar herramientas como Postman o cURL para poder realizar las requests. Los endpoints son los siguientes:</p>

<li><strong>Consultar Inventario: '(GET)' </strong><link>http://127.0.0.1:8080/inventory</link></li>
<li><strong>Registrar Venta: '(POST)' </strong><link>http://127.0.0.1:8080/newSale</link></li>
<li><strong>Generar Reporte: '(POST)' </strong><link>http://127.0.0.1:8080/genReport</link></li>
<br>
<p>Los endpoints de <strong>Registrar Venta</strong> y <strong>Generar Reporte</strong> requieren del uso de un body. El formato de estos se encuentra en el archivo de <strong>'globals.go'</strong> con los nombres de <strong>'SaleRequest'</strong> y <strong>'ReportRequest'</strong> respectivamente.</p>

<h3 align="center"> 🥝 Gracias por usar FRUITstice League 🥝</h3>
