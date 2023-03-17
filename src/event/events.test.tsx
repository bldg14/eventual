import React from "react";
import { render, screen } from "@testing-library/react";
import Event from "./event";

test("renders the Event component", () => {
  const title = "Event";
  const description = "This is the Event Component";

  render(<Event title={title} description={description} />);

  const elTitle = screen.getByText(title);
  const elDesc = screen.getByText(description);
  expect(elTitle).toBeInTheDocument();
  expect(elDesc).toBeInTheDocument();
});
