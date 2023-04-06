import React from "react";
import { render, screen } from "@testing-library/react";
import Event from "./event";

describe("Event", () => {
  const props = {
    title: "Test Title",
    description: "Test Description",
    start: new Date("2023-04-05T14:30:00Z"),
    end: new Date("2023-04-05T16:30:00Z"),
    url: "https://example.com",
    email: "test@example.com",
  };

  it("renders the component with correct props", () => {
    render(<Event {...props} />);

    expect(screen.getByText(props.title)).toBeInTheDocument();
    expect(screen.getByText(props.description)).toBeInTheDocument();
    expect(screen.getByText("9:30 AM - 11:30 AM")).toBeInTheDocument();
    expect(screen.getByText(props.url)).toBeInTheDocument();
    expect(screen.getByText(props.email)).toBeInTheDocument();
  });
});
