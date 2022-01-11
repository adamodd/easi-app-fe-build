import React from 'react';
import { MemoryRouter } from 'react-router-dom';
import { render } from '@testing-library/react';

import { mockSystemInfo } from 'views/Sandbox/mockSystemData';

import BookmarkCard from './index';

describe('BookmarkCard', () => {
  it('matches the snapshot', () => {
    const { asFragment } = render(
      <MemoryRouter>
        <BookmarkCard type="systemList" {...mockSystemInfo[0]} />
      </MemoryRouter>
    );
    expect(asFragment()).toMatchSnapshot();
  });

  it('renders translated headings', () => {
    const { getByText } = render(
      <MemoryRouter>
        <BookmarkCard type="systemList" {...mockSystemInfo[0]} />
      </MemoryRouter>
    );

    // TODO Update expected text output when translations/headings of systemList get solidifed
    expect(getByText('Happiness Achievement Module')).toBeInTheDocument();
    expect(getByText('Officer of Department')).toBeInTheDocument();
    expect(getByText('ATO Status')).toBeInTheDocument();
  });

  it('renders corresponding success health icon for status', () => {
    const { getByTestId } = render(
      <MemoryRouter>
        <BookmarkCard type="systemList" {...mockSystemInfo[0]} />
      </MemoryRouter>
    );

    expect(getByTestId('system-health-icon')).toHaveClass(
      'system-health-icon-success'
    );
  });

  it('renders corresponding warning health icon for status', () => {
    const { getByTestId } = render(
      <MemoryRouter>
        <BookmarkCard type="systemList" {...mockSystemInfo[1]} />
      </MemoryRouter>
    );

    expect(getByTestId('system-health-icon')).toHaveClass(
      'system-health-icon-warning'
    );
  });

  it('renders corresponding fail health icon for status', () => {
    const { getByTestId } = render(
      <MemoryRouter>
        <BookmarkCard type="systemList" {...mockSystemInfo[2]} />
      </MemoryRouter>
    );

    expect(getByTestId('system-health-icon')).toHaveClass(
      'system-health-icon-fail'
    );
  });
});