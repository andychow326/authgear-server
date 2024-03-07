import { useQuery } from "@apollo/client";
import React, {
  createContext,
  useCallback,
  useContext,
  useMemo,
  useState,
} from "react";
import cn from "classnames";
import { Text, ISearchBoxProps, SearchBox, MessageBar } from "@fluentui/react";
import {
  RolesListQueryDocument,
  RolesListQueryQuery,
  RolesListQueryQueryVariables,
} from "./query/rolesListQuery.generated";
import styles from "./RolesScreen.module.css";
import { encodeOffsetToCursor } from "../../util/pagination";
import ScreenContent from "../../ScreenContent";
import NavBreadcrumb from "../../NavBreadcrumb";
import { Context, FormattedMessage } from "@oursky/react-messageformat";
import iconBadge from "../../images/badge.svg";
import PrimaryButton from "../../PrimaryButton";
import { useLocation } from "react-router-dom";
import RolesList from "./RolesList";

const pageSize = 10;

interface LocalSearchBoxProps {
  className?: ISearchBoxProps["className"];
  placeholder?: ISearchBoxProps["placeholder"];
  value?: ISearchBoxProps["value"];
  onChange?: ISearchBoxProps["onChange"];
  onClear?: ISearchBoxProps["onClear"];
}

const LocalSearchBoxContext = createContext<LocalSearchBoxProps | null>(null);

function LocalSearchBox() {
  const value = useContext(LocalSearchBoxContext);
  if (value == null) {
    return null;
  }
  return <SearchBox {...value} />;
}

interface RolesScreenEmptyStateProps {
  className?: string;
}

const RolesScreenEmptyState: React.VFC<RolesScreenEmptyStateProps> =
  function RolesScreenEmptyState(props) {
    const { className } = props;
    const location = useLocation();
    return (
      <div className={cn(className, styles.emptyStateContainer)}>
        <img className={styles.emptyStateIcon} src={iconBadge} />
        <Text className={styles.emptyStateTitle}>
          <FormattedMessage id="RolesScreen.empty-state.title" />
        </Text>
        <Text className={styles.emptyStateSubtitle}>
          <FormattedMessage id="RolesScreen.empty-state.subtitle" />
        </Text>
        <PrimaryButton
          href={`${location.pathname}/add-role`}
          className={styles.emptyStateButton}
          text={<FormattedMessage id={"RolesScreen.empty-state.button"} />}
          iconProps={{ iconName: "Add" }}
        />
      </div>
    );
  };

const RolesScreen: React.VFC = function RolesScreen() {
  const { renderToString } = useContext(Context);
  const [searchKeyword, setSearchKeyword] = useState("");

  const isSearch = searchKeyword !== "";

  const [offset, setOffset] = useState(0);
  // after: is exclusive so if we pass it "offset:0",
  // The first item is excluded.
  // Therefore we have adjust it by -1.
  const cursor = useMemo(() => {
    if (isSearch) {
      // Search always query all rows.
      return null;
    }
    if (offset === 0) {
      return null;
    }
    return encodeOffsetToCursor(offset - 1);
  }, [isSearch, offset]);

  const onChangeSearchKeyword = useCallback((_e, value) => {
    if (value != null) {
      setSearchKeyword(value);
      // Reset offset when search keyword was changed.
      setOffset(0);
    }
  }, []);

  const onClearSearchKeyword = useCallback((_e) => {
    setSearchKeyword("");
    // Reset offset when search keyword was changed.
    setOffset(0);
  }, []);

  const localSearchBoxProps: LocalSearchBoxProps = useMemo(() => {
    return {
      className: styles.searchBox,
      placeholder: renderToString("search"),
      value: searchKeyword,
      onChange: onChangeSearchKeyword,
      onClear: onClearSearchKeyword,
    };
  }, [
    renderToString,
    searchKeyword,
    onChangeSearchKeyword,
    onClearSearchKeyword,
  ]);

  const { data, loading, previousData } = useQuery<
    RolesListQueryQuery,
    RolesListQueryQueryVariables
  >(RolesListQueryDocument, {
    variables: {
      searchKeyword,
      pageSize,
      cursor,
    },
    fetchPolicy: "network-only",
  });

  const isInitialLoading = loading && previousData == null;

  const items = useMemo(() => {
    return [{ to: ".", label: <FormattedMessage id="RolesScreen.title" /> }];
  }, []);

  const isEmpty = !isInitialLoading && (data?.roles?.edges?.length ?? 0) === 0;

  return (
    <LocalSearchBoxContext.Provider value={localSearchBoxProps}>
      <ScreenContent className={styles.content} layout="list">
        <div className={styles.widget}>
          <NavBreadcrumb className="block" items={items} />
        </div>
        <LocalSearchBox />
        {isEmpty ? (
          isSearch ? (
            <MessageBar className={cn(styles.widget, styles.message)}>
              <FormattedMessage id="UsersList.empty.search" />
            </MessageBar>
          ) : (
            <RolesScreenEmptyState className={styles.widget} />
          )
        ) : (
          <RolesList
            className={styles.widget}
            isSearch={isSearch}
            loading={isInitialLoading}
            offset={offset}
            pageSize={pageSize}
            roles={data?.roles ?? null}
          />
        )}
      </ScreenContent>
    </LocalSearchBoxContext.Provider>
  );
};

export default RolesScreen;
