import React from 'react';
import Select from 'react-select';

export const DarklyReactSelect = (props) => (
  <Select
    {...props}
    theme={theme => ({
      ...theme,
      borderRadius: 0,
      colors: {
        ...theme.colors,
        neutral0: 'hsl(0, 0%, 30%)',
        primary25: 'black',
        neutral80: `hsl(0, 0%, ${props.isMulti ? 20 : 100}%)`
      },
    })}
  />
);