import React from "react";
import "./Message-Style.css";
import PseudoDateParser from "@pseudomera/date-parser/dist/dateParser";

const Message = ({ message, imSender, isInfoMessage, createdAt, userName }) => {
  const parsedNumber = PseudoDateParser(createdAt);

  return (
    <div
      className={`message ${imSender === true ? "message--me" : null} ${
        isInfoMessage === true ? "message--info" : null
      } `}
    >
      {imSender === false ? (
        <span className="message__username">{userName} says</span>
      ) : null}
      <span className="message__text">{message}</span>
      <span className="message__time">
        {parsedNumber.day} {parsedNumber.monthShort} {parsedNumber.year}{" "}
        {parsedNumber.hour}:{parsedNumber.minutes} {parsedNumber.amOrPm} 🕒
      </span>
    </div>
  );
};
export default Message;
