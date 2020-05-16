import React from "react";
import DatePicker from "react-datepicker";
import "../../../node_modules/react-datepicker/dist/react-datepicker.css";

class DateRangePanel extends React.Component {
  render() {
    return (
      <div className="date-range-panel">
        <label htmlFor="start">From</label>
        <DatePicker
          name="start"
          selected={this.props.start}
          onChange={val => this.props.setStart(val)}
        />
        <label htmlFor="end">To</label>
        <DatePicker
          name="end"
          selected={this.props.end}
          onChange={val => this.props.setEnd(val)}
        />
      </div>
    );
  }
}

export default DateRangePanel;
