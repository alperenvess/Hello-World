import numpy as np
import pandas as pd
import matplotlib.pyplot as plt
import plotly.express as px
import seaborn as sns
from sklearn.linear_model import LinearRegression
from sklearn.ensemble import RandomForestRegressor
from sklearn.preprocessing import StandardScaler
from sklearn.neural_network import MLPRegressor
from sklearn.model_selection import train_test_split
from sklearn.metrics import r2_score

import os
for dirname, _, filenames in os.walk('/Users/alperenvanlioglu/Desktop/Study/WSB/Projects/World_happiness_reports/2019.csv'):
    for filename in filenames:
        print(os.path.join(dirname, filename))

# Viewing Data
df = pd.read_csv('/Users/alperenvanlioglu/Desktop/Study/WSB/Projects/World_happiness_reports/2019.csv')

df.head()
df.tail()

# Check for missing data
missing_data = df.isnull().sum()

# Print the result
print(missing_data)

df.describe()

# Organizing Contry or Region into Continent Category 
continent_dict = { 
    'Asia': ['Afghanistan', 'Bahrain', 'Iran', 'Iraq', 'Israel', 'Jordan', 'Kuwait', 'Lebanon', 'Oman', 'Pakistan', 'Qatar', 'Saudi Arabia', 'Syria', 'Turkey', 'United Arab Emirates', 'Yemen','Taiwan', 'Singapore', 'Uzbekistan', 'Thailand', 'South Korea', 'Japan', 'Philippines', 'Tajikistan', 'Hong Kong', 'Malaysia', 'Mongolia', 'Kyrgyzstan', 'Turkmenistan', 'Azerbaijan', 'Indonesia', 'China', 'Vietnam', 'Bhutan', 'Nepal', 'Laos', 'Cambodia', 'Palestinian Territories', 'Armenia', 'Georgia', 'Bangladesh', 'Sri Lanka', 'Myanmar', 'India'],
    'Africa': ['Algeria', 'Angola', 'Benin', 'Botswana', 'Burkina Faso', 'Burundi', 'Cabo Verde', 'Cameroon', 'Central African Republic', 'Chad', 'Comoros', 'Congo (Brazzaville)','Congo (Kinshasa)', 'Democratic Republic of the Congo', 'Republic of the Congo', 'Cote d\'Ivoire', 'Djibouti', 'Egypt', 'Equatorial Guinea', 'Eritrea', 'Eswatini', 'Ethiopia', 'Gabon', 'Gambia', 'Ghana', 'Guinea', 'Guinea-Bissau','Ivory Coast', 'Kenya', 'Lesotho', 'Liberia', 'Libya', 'Madagascar', 'Malawi', 'Mali', 'Mauritania', 'Mauritius', 'Morocco', 'Mozambique', 'Namibia', 'Niger', 'Nigeria', 'Rwanda', 'Sao Tome and Principe', 'Senegal', 'Seychelles', 'Sierra Leone', 'Somalia', 'South Africa', 'South Sudan', 'Sudan','Swaziland', 'Tanzania', 'Togo', 'Tunisia', 'Uganda', 'Zambia', 'Zimbabwe'],
    'North America': ['Antigua and Barbuda', 'Bahamas', 'Barbados', 'Belize', 'Canada', 'Costa Rica', 'Cuba', 'Dominica', 'Dominican Republic', 'El Salvador', 'Grenada', 'Guatemala', 'Haiti', 'Honduras', 'Jamaica', 'Mexico', 'Nicaragua', 'Panama', 'Saint Kitts and Nevis', 'Saint Lucia', 'Saint Vincent and the Grenadines', 'Trinidad and Tobago', 'United States'],
    'South America': ['Argentina', 'Bolivia', 'Brazil', 'Chile', 'Colombia', 'Ecuador', 'Guyana', 'Paraguay', 'Peru', 'Suriname', 'Trinidad & Tobago', 'Uruguay', 'Venezuela'],
    'Europe': ['Albania', 'Andorra', 'Austria','Azerbaijan', 'Belarus', 'Belgium', 'Bosnia and Herzegovina', 'Bulgaria', 'Croatia', 'Cyprus', 'Czech Republic', 'Denmark', 'Estonia', 'Finland', 'France', 'Germany', 'Greece','Georgia', 'Hungary', 'Iceland', 'Ireland', 'Italy', 'Kazakhstan', 'Kosovo', 'Latvia', 'Liechtenstein', 'Lithuania', 'Luxembourg', 'Malta', 'Moldova', 'Monaco', 'Montenegro', 'Netherlands','Northern Cyprus', 'North Macedonia', 'Norway', 'Poland', 'Portugal', 'Romania', 'Russia', 'San Marino', 'Serbia', 'Slovakia', 'Slovenia', 'Spain', 'Sweden', 'Switzerland', 'Ukraine', 'United Kingdom', 'Vatican City'],
    'Oceania': ['Australia', 'Fiji', 'Kiribati', 'Marshall Islands', 'Micronesia', 'Nauru', 'New Zealand', 'Palau', 'Papua New Guinea', 'Samoa', 'Solomon Islands', 'Tonga', 'Tuvalu', 'Vanuatu']
    }

df['Continent'] = df['Country or region'].apply(lambda x: next((k for k in continent_dict if x in continent_dict[k]), None))

df.head()
df.tail()

# Check for missing data
missing_data = df.isnull().sum()

# Print the result
print(missing_data)

# Check for duplicates
duplicates = df[df.duplicated()]

# Print the duplicates
print(duplicates)

df = df.sort_values('Continent')

# Create a violin plot
fig_violin = px.violin(df, x='Continent', y='Score', color='Continent',
                       title='Happiness Scores by Continent (Violin Plot)',
                       labels={'Continent': 'Continent', 'Score': 'Happiness Score'})
fig_violin.show()

# Group the data by continent and calculate the mean and median scores
grouped = df.groupby('Continent')['Score']
mean_scores = grouped.mean()
median_scores = grouped.median()

# Print the results
print('Mean scores by continent:')
print(mean_scores)
print('\nMedian scores by continent:')
print(median_scores)

# Create scatter plots with regression lines for each continent
scatter_plots = px.scatter(df, x="Score", y="GDP per capita", color="Continent", trendline="ols", facet_col="Continent")
scatter_plots.update_layout(title_text='GDP per capita vs. Happiness Score', xaxis_title='Happiness Score')
scatter_plots.update_xaxes(title='Happiness Score')
scatter_plots.for_each_annotation(lambda x: x.update(text=x.text.split("=")[-1]))
scatter_plots.show()

scatter_plots = px.scatter(df, x="Score", y="Social support", color="Continent", trendline="ols", facet_col="Continent")
scatter_plots.update_layout(title_text='Social support vs. Happiness Score', xaxis_title='Happiness Score')
scatter_plots.update_xaxes(title='Happiness Score')
scatter_plots.for_each_annotation(lambda x: x.update(text=x.text.split("=")[-1]))
scatter_plots.show()

scatter_plots = px.scatter(df, x="Score", y="Healthy life expectancy", color="Continent", trendline="ols", facet_col="Continent")
scatter_plots.update_xaxes(title='Happiness Score')
scatter_plots.update_layout(title_text='Healthy life expectancy vs. Happiness Score', xaxis_title='Happiness Score')
scatter_plots.for_each_annotation(lambda x: x.update(text=x.text.split("=")[-1]))
scatter_plots.show()

scatter_plots = px.scatter(df, x="Score", y="Freedom to make life choices", color="Continent", trendline="ols", facet_col="Continent")
scatter_plots.update_xaxes(title='Happiness Score')
scatter_plots.update_layout(title_text='Freedom to make life choices vs. Happiness Score', xaxis_title='Happiness Score')
scatter_plots.for_each_annotation(lambda x: x.update(text=x.text.split("=")[-1]))
scatter_plots.show()

scatter_plots = px.scatter(df, x="Score", y="Generosity", color="Continent", trendline="ols", facet_col="Continent")
scatter_plots.update_layout(title_text='Generosity vs. Happiness Score', xaxis_title='Happiness Score')
scatter_plots.update_xaxes(title='Happiness Score')
scatter_plots.for_each_annotation(lambda x: x.update(text=x.text.split("=")[-1]))
scatter_plots.show()

scatter_plots = px.scatter(df, x="Score", y="Perceptions of corruption", color="Continent", trendline="ols", facet_col="Continent")
scatter_plots.update_layout(title_text='Perceptions of corruption vs. Happiness Score', xaxis_title='Happiness Score')
scatter_plots.update_xaxes(title='Happiness Score')
scatter_plots.for_each_annotation(lambda x: x.update(text=x.text.split("=")[-1]))
scatter_plots.show()

# Convert the categorical variable 'Continent' to a numeric variable
df['Continent'] = pd.factorize(df['Continent'])[0]

df.head()

# Drop the 'Overall rank' column 
# Since Rank and Happiness Score is negatively correlated by nature
dfcor = df.drop(columns=['Overall rank'])

# Compute the correlation matrix for numeric columns
corr_matrix = dfcor.corr(numeric_only=True)

# Create a heatmap using Seaborn
sns.heatmap(corr_matrix, annot=True, cmap='coolwarm')

# Set the title of the plot
plt.title('Correlation Heatmap')

# Show the plot
plt.show()

# Split the data into training and testing sets
X = df[['GDP per capita', 'Social support', 'Healthy life expectancy', 'Freedom to make life choices', 'Generosity', 'Perceptions of corruption', 'Continent']]
y = df['Score']
X_train, X_test, y_train, y_test = train_test_split(X, y, test_size=0.3, random_state=42)

# Create a linear regression model and fit it to the training data
model = LinearRegression()
model.fit(X_train, y_train)

# Predict the target variable for the testing data
y_pred = model.predict(X_test)

# Calculate the R-squared value of the model
r2 = r2_score(y_test, y_pred)

# Plot the predicted score against the actual score
plt.scatter(y_test, y_pred)
plt.plot([y_test.min(), y_test.max()], [y_test.min(), y_test.max()], 'k--', lw=4)
plt.xlabel('Actual Score')
plt.ylabel('Predicted Score')
plt.title('Predicted Score vs Actual Score')
plt.show()

# Print a summary of the linear regression model
print('Linear Regression Model Summary')
print('------------------')
print('Intercept:', model.intercept_)
for i, col in enumerate(X.columns):
    print(col, ':', model.coef_[i])
print('------------------')
print('R-squared:', r2)
print('------------------')

# Split the data into training and testing sets
X = df[['GDP per capita', 'Social support', 'Healthy life expectancy', 'Freedom to make life choices', 'Generosity', 'Perceptions of corruption', 'Continent']]
y = df['Score']
X_train, X_test, y_train, y_test = train_test_split(X, y, test_size=0.3, random_state=42)

# Create a random forest model and fit it to the training data
model = RandomForestRegressor(n_estimators=500, random_state=42)
model.fit(X_train, y_train)

# Predict the target variable for the testing data
y_pred = model.predict(X_test)

# Calculate the R-squared value of the model
r2 = r2_score(y_test, y_pred)

# Plot the predicted score against the actual score
plt.scatter(y_test, y_pred)
plt.plot([y_test.min(), y_test.max()], [y_test.min(), y_test.max()], 'k--', lw=4)
plt.xlabel('Actual Score')
plt.ylabel('Predicted Score')
plt.title('Predicted Score vs Actual Score')
plt.show()

# Print a summary of the random forest model
print('Random Forest Model Summary')
print('------------------')
for i, col in enumerate(X.columns):
    print(col, ':', model.feature_importances_[i])
print('------------------')
print('R-squared:', r2)
print('------------------')