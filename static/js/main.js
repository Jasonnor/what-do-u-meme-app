$(
    (function() {
      "use strict";
      // Search Text
      let userInput = $("#userInput");
      // Search Btn
      let searchBtn = $("#searchBtn");
      // Trend Btn
      let trendBtn = $("#trendBtn");
      // More Images btn
      let moreImagesBtn = $("#moreImagesBtn");
      // Result Image area
      let imageContainer = $(".image-container");
      // Count of returned results
    //   let perPageCount = 20;
      // variable for current page requested so that when the user scrolls, it is increamented and json is requested
      let currentPage = 1;
      let n_result = 30; 
  
      // SEARCH BTN EVENT LISTENER
      searchBtn.on("click", function() {
        // Clear existing images
        // imageContainer.html("");
        validateInput();
      }); // end of SEARCH BTN EVENT LISTENER
  
      //USER SEARCH SUBMIT [ENTER BTN] EVENT LISTENER
      userInput.on("keyup", function(e) {
        if (e.keyCode == 13) {
          // Clear exisitng images
          // imageContainer.html("");
          validateInput();
        }
      }); // end of USER SEARCH SUBMIT [ENTER BTN] EVENT LISTENER

      // TREND BTN EVENT LISTENER
      trendBtn.on("click", function() {
        imageContainer.html("");
        getTrendingImageData();
      });
  
    //   // MORE IMAGES BTN
    //   moreImagesBtn.on("click", function() {
    //     currentPage++;
    //     console.log(currentPage);
    //     validateInput();
    //   }); // END OF MORE IMAGES BTN
  
      // VALIDATE USER INPUT
      function validateInput() {
        //check if query len is valid
        if (userInput.val().length != 0) {
          // Call the AJAX req passing in the user search
          imageContainer.html("");
          getImageData(userInput.val());
        } else {
          alert("Please fill in text on input bar!")
          // userInput.focus();
        } // end of validation
      } //end of VALIDATE USER INPUT
  
      // GET JSON FROM FLICKR API
      function getImageData(query) {
        $.ajax({
          type: "GET",
          url: `/mock/search_by_text?input=${query}&n_result=${n_result}&page=${currentPage}`,
          success: function(data) {
            // Call the handle Data function and pass the response
            data = JSON.parse(data);
            handleData(data);
          },
          error: function(e) {
            alert("Failed to load data from API");
            // moreImagesBtn.html("Something Went Wrong");
            console.log(e);
          },
          beforeSend: function() { $('.loader').show();},
          complete: setTimeout(function() { $('.loader').hide(); }, 1000)
        });
      } // END OF getImageData

      function getTrendingImageData() {
        $.ajax({
          type: "GET",
          url: `http://localhost:8080/mock/get_trending?n_result=${n_result}&page=${currentPage}`,
          success: function(data) {
            // Call the handle Data function and pass the response
            data = JSON.parse(data);
            handleData(data);
          },
          error: function(e) {
            alert("Failed to load data from API");
            // moreImagesBtn.html("Something Went Wrong");
            console.log(e);
          },
          beforeSend: function() { $('.loader').show();},
          complete: setTimeout(function() { $('.loader').hide(); }, 1000)
        });
      } // END OF getImageData
  
      // Build urls for the images
      function handleData(data) {
        // console.log(data);
        data.forEach(function(currentPhoto, index, array) {
          // Image URL > Will be built and made an link
          let photoURL = currentPhoto.image_url;
          // alt text
          let title = currentPhoto.title;
          // plug the data to page
          pushImages(photoURL, title);
        });
        //show the more btn
        // moreImagesBtn.slideDown();
      } // END OF handleData
  
      // Built HTML template and push the images to the webpage
      function pushImages(url, title) {
        // Build the HTML element
        let htmlText = `<img alt="${title}" src="${url}">`;
        imageContainer.append(htmlText);
        // Add the Materialize functionality to the images
        // $(".materialboxed").materialbox();
      } // END OF pushImages

      trendBtn.trigger("click");

    })() // end of self self executing anonymous function
  );
  
