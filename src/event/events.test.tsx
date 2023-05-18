import React from "react";
import { render, screen } from "@testing-library/react";
import { Event, formatTime } from "./event";

describe("Event", () => {
  const props = {
    title: "Test Title",
    description: "Test Description",
    start: new Date(),
    end: new Date(Date.now() + 1000 * 60 * 60),
    url: "https://example.com",
    email: "test@example.com",
  };

  it("renders the component with correct props", () => {
    render(<Event {...props} />);

    const expectTime = `${formatTime(props.start)} - ${formatTime(props.end)}`

    expect(screen.getByText(props.title)).toBeInTheDocument();
    expect(screen.getByText(props.description)).toBeInTheDocument();
    expect(screen.getByText(expectTime)).toBeInTheDocument();
    expect(screen.getByText(props.url)).toBeInTheDocument();
    expect(screen.getByText(props.email)).toBeInTheDocument();
  });
});
