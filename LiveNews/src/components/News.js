import { React, useState, useEffect, Fragment } from "react";
import Item from "./Item";
import Spinner from "./Spinner";
import InfiniteScroll from "react-infinite-scroll-component";
import { getNews } from "../utils/newsApi";
import { getWeather } from "../utils/weatherApi";

const News = ({ category }) => {
  const [articles, setArticles] = useState([]);
  const [loading, setLoading] = useState(false);
  const [total, setTotal] = useState(5);
  const [resp, setResp] = useState("");
  const [cityName, setCity] = useState("");
  const [page, setPage] = useState(1);
  const pageSize = 6;

  const fetchData = async () => {
    setLoading(true);
    if (category !== "Weather") {
      setPage(page + 1);
      const response = await getNews(category, page, pageSize);
      if (response?.articles) {
        setArticles(articles?.concat(response?.articles));
        setTotal(total + response?.articles?.length);
      }
    } else if (category === "Weather") {
      const res = await getWeather(cityName);
      setResp(res?.data);
    }
    setLoading(false);
  };

  useEffect(() => {
    fetchData(category);
  }, [category]);

  return (
    <>
      <h1 className="text-center my-4">
        Top News On <span className="badge bg-dark">{category}</span>
      </h1>
      {loading && <Spinner />}
      {articles.length !== 0 && (
        <Fragment>
          <div className="container my-5">
            <InfiniteScroll
              dataLength={articles?.length}
              next={fetchData}
              hasMore={articles?.length !== total}
              loader={loading && <Spinner />}
            >
              <div className="row my-2">
                {articles?.map((element, index) => (
                  <div className="col-4" key={index}>
                    <Item
                      title={element?.title || ""}
                      imgUrl={element?.UrlToImage}
                      description={element?.description || "No Content"}
                      url={element?.url || ""}
                      author={element?.author || "Unknown"}
                      date={element?.publishedAt || "00:00:00"}
                      source={element?.source?.Name || ""}
                      category={category}
                    />
                  </div>
                ))}
              </div>
            </InfiniteScroll>
          </div>
        </Fragment>
      )}
      {resp && (
        <Fragment>
          <div className="container my-5">
            <div className="row my-2 justify-content-center">
              <div className="col-4">
                <div className="card">
                  <img
                    src={"static/bg1.png"}
                    className="card-img-top"
                    alt="..."
                  />
                  <div className="card-body m-auto">
                    <h5 className="card-title">
                      {resp?.name}
                      <span style={{ left: "90%", zIndex: "1" }}>
                        - {resp?.sys?.country}
                      </span>
                    </h5>
                    <p className="card-text">{Math.round(resp?.main?.temp - 273.15)} Celcius</p>
                    <p className="card-text">
                      latestBy - {resp.Date.year}/{resp.Date.month}/{resp.Date.day}
                    </p>
                    <input
                      className="city-ip"
                      type="text"
                      placeholder="Enter your city name - lucknow"
                      onChange={(e) => setCity(e.target.value)}
                    />&nbsp;&nbsp;&nbsp;
                    <input
                      type="button"
                      value="Submit"
                      className="btn-primary"
                      onClick={() => {
                        fetchData("Weather");
                      }}
                    />
                  </div>
                </div>
              </div>
            </div>
          </div>
        </Fragment>
      )}
    </>
  );
};

export default News;
