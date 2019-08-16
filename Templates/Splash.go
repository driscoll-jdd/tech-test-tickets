package Templates

import (
	"strings"
)

func Page_Splash() string {

		output := strings.TrimSpace(`

<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">

  <title>Book Your Tickets</title>

  <link href="https://cdn.driscoll.co/css/robust.css" rel="stylesheet">
</head>

<body class="bg-light">

<style>

/**
 * The CSS shown here will not be introduced in the Quickstart guide, but shows
 * how you can use CSS to style your Element's container.
 */
.StripeElement {
  box-sizing: border-box;

  height: 40px;

  padding: 10px 12px;

  border: 1px solid transparent;
  border-radius: 4px;
  background-color: white;

  box-shadow: 0 1px 3px 0 #e6ebf1;
  -webkit-transition: box-shadow 150ms ease;
  transition: box-shadow 150ms ease;
}

.StripeElement--focus {
  box-shadow: 0 1px 3px 0 #cfd7df;
}

.StripeElement--invalid {
  border-color: #fa755a;
}

.StripeElement--webkit-autofill {
  background-color: #fefde5 !important;
}

</style>

  <nav class="navbar navbar-lg navbar-expand-lg navbar-dark bg-primary">
    <div class="container">
      <a class="navbar-brand" href="/">Book Your Tickets</a>
      <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarCollapse" aria-controls="navbarCollapse" aria-expanded="false" aria-label="Toggle navigation">
        <span class="navbar-toggler-icon"></span>
      </button>

      <div class="collapse navbar-collapse" id="navbarCollapse">

      </div>
    </div>
  </nav>

  <main class="main" role="main">

	<div class="container h-100">
      <div class="row align-items-center h-100">
        <div class="col-sm-10 col-md-8 col-lg-6 mx-auto my-4">

          <div class="text-center">
            <h1 class="h3">Exclusive Event</h1>
            <p class="lead">
              Enter your name to book your ticket.
            </p>
			<p class="text-warning">
				Limit: One per customer.
			</p>
          </div>

			<div class="alert alert-danger" role="alert" id="faultBox" style="display: none;"></div>
			<div class="alert alert-success" role="alert" id="payBox" style="display: none;"></div>

          <div class="card" id="bookCard">
            <div class="card-body">
              <div class="m-sm-4">
				  <div class="form-group">
					<label>Your Details</label>
					<input class="form-control form-control-lg" type="text" id="name" name="name" placeholder="Enter your name" />
				  </div>
				  <div class="text-center mt-3">
					<button type="submit" class="btn btn-lg btn-primary" id="bookButton" onclick="reserveTicket();">Book Your Ticket</button>
				  </div>
              </div>
            </div>
          </div><!-- /.card -->

		<div class="card" id="timeCard" style="display: none;">
            <div class="card-body">

				<h6>Purchase Your Ticket</h6>

				<p>
					Your ticket will be released if you do not purchase it in time.
				</p>

				<div class="alert alert-warning" id="timeBox" role="alert"></div>

			</div>
		</div>


        </div>
      </div><!-- /.row -->
    </div>


</main>


	<script src="https://cdn.driscoll.co/js/bundle.js"></script>

<script>

	function reserveTicket() {

		var ticketIO;
  		ticketIO=new XMLHttpRequest();
  		ticketIO.onreadystatechange = function() {
    
			if (this.readyState == 4 && this.status == 200) {
      
				response = JSON.parse(this.responseText);

				if(response.Outcome) {

					box = document.getElementById('payBox');
					box.innerHTML = "Tickets reserved"

					countdown(response.Ticket.ReleasePoint);

					document.getElementById('payBox').style.display = 'block';
					document.getElementById('timeCard').style.display = 'block';
					document.getElementById('bookCard').style.display = 'none';
					document.getElementById('bookButton').style.display = 'none';
					document.getElementById('faultBox').style.display = 'none';
				
				} else {

					if(response.Faults.Name) {

						box = document.getElementById('faultBox');
						box.innerHTML = "You must enter your name to book a ticket";
						box.style.display = "block";

						document.getElementById('name').focus();
					}
					else
					{
						box = document.getElementById('faultBox');
						box.innerHTML = response.Error;
						box.style.display = "block";

						document.getElementById('name').focus();
					}
				}
    		}
  		};

		ticketIO.open("GET", "/reserve/" + document.getElementById('name').value, true);
		ticketIO.send();
	}

	function countdown(date) {

		// Set the date we're counting down to
var countDownDate = new Date(date).getTime();

// Update the count down every 1 second
var x = setInterval(function() {

  // Get today's date and time
  var now = new Date().getTime();

  // Find the distance between now and the count down date
  var distance = countDownDate - now;

  // Time calculations for days, hours, minutes and seconds
  var days = Math.floor(distance / (1000 * 60 * 60 * 24));
  var hours = Math.floor((distance % (1000 * 60 * 60 * 24)) / (1000 * 60 * 60));
  var minutes = Math.floor((distance % (1000 * 60 * 60)) / (1000 * 60));
  var seconds = Math.floor((distance % (1000 * 60)) / 1000);

  // Display the result in the element with id="demo"
  document.getElementById("timeBox").innerHTML = minutes + "m " + seconds + "s ";

  // If the count down is finished, write some text 
  if (distance < 0) {
    clearInterval(x);
    document.getElementById("timeBox").innerHTML = "Your ticket has now been released. Please re-book.";
  }
}, 1000);

	}

</script>

</body>
</html>

		`)

		return output
	}
