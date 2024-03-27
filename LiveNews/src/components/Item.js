import React, { Component } from "react";

export default class Item extends Component {
  render() {
    const { title, description, url, imgUrl, author, date, source } =
    this.props;
    return (
      <div className="card">
        <img src={imgUrl} className="card-img-top" alt="..." />
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
}
