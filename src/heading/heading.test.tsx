import React from "react";
import { render, screen } from "@testing-library/react";
import Heading from "./heading";

test("renders the Heading component", () => {
  const title = "sup";
  render(<Heading>{title}</Heading>);
  const el = screen.getByText(title);
  expect(el).toBeInTheDocument();
});
