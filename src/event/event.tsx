import Container from "react-bootstrap/Container";
import Row from "react-bootstrap/Row";
import Col from "react-bootstrap/Col";

interface EventProps {
  title: string;
  description: string;
  start: Date;
  end: Date;
  url: string;
  email: string;
}

export const formatTime = (time: Date): string => {
  return new Intl.DateTimeFormat("en-US", {
    hour: "numeric",
    minute: "numeric",
    hour12: true,
  }).format(time);
};

export const Event = ({
  title,
  description,
  start,
  end,
  url,
  email,
}: EventProps) => {
  return (
    <Container className="rounded overflow-hidden bg-info-subtle">
      <Row>
        <Col>
          <h3>{title}</h3>
        </Col>
        <Col>
          <p className="text-end">
            {formatTime(start)} - {formatTime(end)}
          </p>
        </Col>
      </Row>
      <Row>
        <p>{description}</p>
      </Row>
      <Row>
        <Col>
          <a className="text-decoration-none link-secondary" href={url}>
            {url}
          </a>
        </Col>
        <Col className="text-end">
          <a
            className="text-decoration-none link-secondary"
            href={"mailto:" + email}
          >
            {email}
          </a>
        </Col>
      </Row>
    </Container>
  );
};
