interface HeadingProps {
  children: string;
}

const Heading = ({ children }: HeadingProps) => {
  return <h1>{children}</h1>;
};

export default Heading;
