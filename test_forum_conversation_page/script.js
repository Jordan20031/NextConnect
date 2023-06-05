const imgDiv = document.querySelector('.profile-pic-div');
const img = document.querySelector('#photo');
const file = document.querySelector('#file');

imgDiv.addEventListener('mouseenter', function() {
  file.style.display = "block";
});

imgDiv.addEventListener('mouseleave', function() {
  file.style.display = "none";
});

file.addEventListener('change', function() {
  const choosedFile = this.files[0];
  if (choosedFile) {
    const reader = new FileReader();
    reader.addEventListener('load', function() {
      img.setAttribute('src', reader.result);
    });
    reader.readAsDataURL(choosedFile);
  }
});

function sendMessage() {
  var userInput = document.getElementById('user-input');
  var messageContainer = document.querySelector('.message-container');

  if (userInput.value.trim() === '') {
    return; 
  }

  var userMessage = createMessageElement(userInput.value, 'user-message');
  messageContainer.appendChild(userMessage);

  // Vérifie s'il y a une image sélectionnée
  var imageInput = document.getElementById('file');
  if (imageInput.files.length > 0) {
    var reader = new FileReader();
    reader.onload = function(event) {
      var imageUrl = event.target.result;
      appendImageToMessage(userMessage, imageUrl);
      clearImageInput(); // Réinitialiser la sélection de l'image
    };
    reader.readAsDataURL(imageInput.files[0]);
  } else {
    userInput.value = ''; // Réinitialiser la valeur de l'entrée utilisateur si aucune image n'est sélectionnée
    clearImageInput(); // Réinitialiser la sélection de l'image
  }
}

function appendImageToMessage(messageElement, imageUrl) {
  var imageElement = document.createElement('img');
  imageElement.src = imageUrl;
  imageElement.className = 'preview';

  var messageInfoElement = document.createElement('div');
  messageInfoElement.className = 'message-info';
  messageInfoElement.innerText = 'Moi - ' + getCurrentTimestamp();

  messageElement.appendChild(imageElement);
  messageElement.appendChild(messageInfoElement);
}

function createMessageElement(messageText, className) {
  var messageElement = document.createElement('div');
  messageElement.className = 'message ' + className;

  var messageTextElement = document.createElement('span');
  messageTextElement.innerText = messageText;

  var messageInfoElement = document.createElement('div');
  messageInfoElement.className = 'message-info';
  messageInfoElement.innerText = 'Moi - ' + getCurrentTimestamp();

  messageElement.appendChild(messageTextElement);
  messageElement.appendChild(messageInfoElement);

  return messageElement;
}

function getCurrentTimestamp() {
  var date = new Date();
  var hours = date.getHours().toString().padStart(2, '0');
  var minutes = date.getMinutes().toString().padStart(2, '0');
  var timestamp = hours + ':' + minutes;

  return timestamp;
}

function clearImageInput() {
  file.value = '';
  img.src = '';
}
