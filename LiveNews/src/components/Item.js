import React, { useEffect, useState } from "react";
import { getImage } from "../utils/imageApi";

const Item = ({title, description, url, imgUrl, author, date, source , category})=> {
  const [img, setImg] = useState(imgUrl);

  const getImg = async ()=>{
        if (img === ""){
        const response = await getImage(category);
        setImg(response?.data?.photos?.[0]?.src?.medium)
      }
  }
  getImg();
    return (
      <div className="card">
        <img src={img} className="card-img-top" alt="..." />
        <div className="card-body">
          <h5 className="card-title">
            {title} - 
            <span className="" style={{ left: "90%", zIndex: "1" }}>
              {source}
            </span>
          </h5>
          <p className="card-text">{description}</p>
          <p className="card-text">
            <small className="text-muted">
              By {author} on {new Date(date).toGMTString()}
            </small>
          </p>
          <a
            href={url}
            target="_blank"
            rel="noreferrer"
            className="btn btn-sm btn-dark"
          >
            Read More?
          </a>
        </div>
      </div>
    );
  }

export default Item;
