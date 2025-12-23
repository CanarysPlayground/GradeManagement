import { test, expect } from '@playwright/test';

test('grades list renders', async ({ page }) => {
  // wait until network is idle so Angular has time to bootstrap and render
  await page.goto('http://localhost:4200', { waitUntil: 'networkidle' });

  // use getByText which is more robust for visible text assertions
  const gradesText = page.getByText('Grades');

  // give a small timeout to allow client rendering
  await expect(gradesText).toBeVisible({ timeout: 5000 });
});
