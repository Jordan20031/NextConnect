const imgDiv = document.querySelector('.profile-pic-div');
const img = document.querySelector(`#photo`);
const file = document.querySelector('#file');
const uploadBtn = document.querySelector('#uploadBtn');
imgDiv.addEventListener('mouseenter', function()
{
    uploadBtn.getElementsByClassName.display = "block";
});
imgDiv.addEventListener('mouseleave', function()
{
    uploadBtn.getElementsByClassName.display = "none";

});
file.addEventListener('change', function(){
    const choosedFile = this.files[0];
    if (choosedFile) {
        const reader = new FileReader();
        reader.addEventListener('load', function
        (){
            img.setAttribute('src',
            this.result);
        });

        reader.readAsDataURL(choosedFile);   
    }
})