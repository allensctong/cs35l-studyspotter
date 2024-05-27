import React, {useState, useEffect} from "react";
import './App.css'

function UploadImages(){
    const[images, setImages]= useState([]);
    const[imageURLs, setImageURLs]=useState([]);

    useEffect(()=> {

        if(images.length<1) return;
        const newImageUrls=[];
        images.forEach(image=> newImageUrls.push(URL.createObjectURL(image)));
        setImageURLs(newImageUrls);

        return()=>{
            newImageUrls.forEach(url => URL.revokeObjectURL(url));
        }; 
    }, [images]);

    function onImageChange(e){
        const files=Array.from(e.target.files);
        const validImageFiles=files.filter(file=> {
            const fileExtension=file.name.split('.').pop().toLowerCase();
            return ['jpg', 'jpeg', 'png'].includes(fileExtension);
        });
        if(validImageFiles.length){
            setImages(validImageFiles);
        }
        

    }

    return(
        <>
            <input type="file" multiple accpet=".jpg,.jpeg,.png" onChange={onImageChange} />
            <div className="image=preview-container">
            {imageURLs.map((imageSrc, index) => (
          <img key={index} src={imageSrc} alt={`uploaded image ${index + 1}`} className="uploaded-image" />
        ))}
            </div>
           
        
        </>


    );

}




export default UploadImages;