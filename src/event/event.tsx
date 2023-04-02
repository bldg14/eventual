interface EventProps {
  title: string;
  description: string;
  start: Date;
  end: Date;
  url: string;
  email: string;
}

const Event = ({ title, description, start, end, url, email }: EventProps) => {
  return (
    <div className="event-container">
      <h2 className="event-title">{title}</h2>
      <div className="event-details">
        <div className="event-time-container">
          <p className="event-time">
            {start.toLocaleTimeString()} - {end.toLocaleTimeString()}
          </p>
        </div>
        <div className="event-description-container">
          <p className="event-description">{description}</p>
        </div>
      </div>
      <div className="event-footer">
        <button className="event-url">{url}</button>
        <button className="event-email">{email}</button>
      </div>
    </div>
  );
};

export default Event;
