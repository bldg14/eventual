import Heading from "../heading/heading";

const Event = ({
  title,
  description,
}: {
  title: string;
  description: string;
}) => {
  return (
    <div>
      <Heading title={title} />
      <p>{description}</p>
    </div>
  );
};

export default Event;
